// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package serviceexporter

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-micro-saas/service-kit/server"
	"github.com/go-micro-saas/service-kit/setup"
	"github.com/go-micro-saas/service-kit/testdata/ping-service/internal/biz/biz"
	"github.com/go-micro-saas/service-kit/testdata/ping-service/internal/data/data"
	"github.com/go-micro-saas/service-kit/testdata/ping-service/internal/service/service"
)

// Injectors from wire.go:

func exportServices(launcherManager setuputil.LauncherManager, hs *http.Server, gs *grpc.Server) (serverutil.ServiceInterface, error) {
	logger, err := setuputil.GetLogger(launcherManager)
	if err != nil {
		return nil, err
	}
	homeService := service.NewHomeService(logger)
	websocketBizRepo := biz.NewWebsocketBiz(logger)
	websocketService := service.NewWebsocketService(logger, websocketBizRepo)
	serviceAPIManager, err := setuputil.GetServiceAPIManager(launcherManager)
	if err != nil {
		return nil, err
	}
	pingDataRepo := data.NewPingData(logger)
	pingBizRepo := biz.NewPingBiz(logger, serviceAPIManager, pingDataRepo)
	srvPingServer := service.NewPingService(logger, pingBizRepo)
	srvTestdataServer := service.NewTestdataService(logger, websocketBizRepo)
	serviceInterface, err := service.RegisterServices(hs, gs, homeService, websocketService, srvPingServer, srvTestdataServer)
	if err != nil {
		return nil, err
	}
	return serviceInterface, nil
}
