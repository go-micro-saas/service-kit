//go:build wireinject
// +build wireinject

package exportservices

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	serverutil "github.com/go-micro-saas/service-kit/server"
	setuputil "github.com/go-micro-saas/service-kit/setup"
	"github.com/go-micro-saas/service-kit/testdata/ping-service/internal/biz/biz"
	"github.com/go-micro-saas/service-kit/testdata/ping-service/internal/data/data"
	"github.com/go-micro-saas/service-kit/testdata/ping-service/internal/service/service"
	"github.com/google/wire"
)

func exportServices(launcherManager setuputil.LauncherManager, hs *http.Server, gs *grpc.Server) (*serverutil.Services, func(), error) {
	panic(wire.Build(
		// service
		setuputil.GetLogger,
		data.NewPingData,
		biz.NewWebsocketBiz, biz.NewPingBiz,
		service.NewHomeService, service.NewWebsocketService,
		service.NewPingService, service.NewTestdataService,
		// register services
		service.RegisterServices,
	))
	return &serverutil.Services{}, func() {}, nil
}
