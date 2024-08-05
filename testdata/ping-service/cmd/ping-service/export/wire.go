//go:build wireinject
// +build wireinject

package pingexport

import (
	serverutil "github.com/go-micro-saas/service-kit/server"
	setuputil "github.com/go-micro-saas/service-kit/setup"
	"github.com/go-micro-saas/service-kit/testdata/ping-service/api"
	"github.com/google/wire"
)

//var (
//	HomeService      = wire.NewSet(setuputil.GetLogger, service.NewHomeService)
//	WebsocketService = wire.NewSet(setuputil.GetLogger, biz.NewWebsocketBiz, service.NewWebsocketService)
//	PingService      = wire.NewSet(setuputil.GetLogger, biz.NewPingBiz, service.NewPingService)
//	TestdataService  = wire.NewSet(setuputil.GetLogger, biz.NewWebsocketBiz, service.NewTestdataService)
//)

func initServices(launcherManager setuputil.LauncherManager) (*Services, error) {
	panic(wire.Build(
		api.GetAuthWhiteList,
		serverutil.NewGRPCServer, serverutil.NewHTTPServer,
		NewServices,
	))
	return nil, nil
}
