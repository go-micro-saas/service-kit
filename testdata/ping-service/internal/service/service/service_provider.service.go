package service

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/mux"
	pingservicev1 "github.com/ikaiguang/go-srv-kit/api/ping/v1/services"
	testdataservicev1 "github.com/ikaiguang/go-srv-kit/api/testdata/v1/services"
	stdlog "log"
)

func RegisterServices(
	hs *http.Server, gs *grpc.Server,
	homeService HomeService,
	websocketService WebsocketService,
	pingService pingservicev1.SrvPingServer,
	testdataService testdataservicev1.SrvTestdataServer,
) error {
	// grpc
	if gs != nil {
		pingservicev1.RegisterSrvPingServer(gs, pingService)
		testdataservicev1.RegisterSrvTestdataServer(gs, testdataService)
	}

	// http
	if hs != nil {
		pingservicev1.RegisterSrvPingHTTPServer(hs, pingService)
		testdataservicev1.RegisterSrvTestdataHTTPServer(hs, testdataService)

		// special
		RegisterSpecialRouters(hs, homeService, websocketService)
	}
	return nil
}

func RegisterSpecialRouters(hs *http.Server, homeService HomeService, websocketService WebsocketService) {
	// router
	router := mux.NewRouter()

	stdlog.Println("|*** 注册路由：Root(/)")
	router.HandleFunc("/", homeService.Homepage)
	hs.Handle("/", router)

	stdlog.Println("|*** 注册路由：Websocket")
	router.HandleFunc("/ws/v1/testdata/websocket", websocketService.TestWebsocket)

	// router
	hs.Handle("/ws/v1/websocket", router)
}
