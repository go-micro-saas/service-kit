//go:build wireinject
// +build wireinject

package runservices

import (
	"github.com/go-kratos/kratos/v2"
	configutil "github.com/go-micro-saas/service-kit/config"
	serverutil "github.com/go-micro-saas/service-kit/server"
	setuputil "github.com/go-micro-saas/service-kit/setup"
	"github.com/go-micro-saas/service-kit/testdata/ping-service/api"
	serviceexport "github.com/go-micro-saas/service-kit/testdata/ping-service/cmd/ping-service/export"
	"github.com/google/wire"
)

func initServiceApp(configFilePath string, configOpts ...configutil.Option) (*kratos.App, func(), error) {
	panic(wire.Build(
		setuputil.NewLauncherManagerWithCleanup,
		api.GetAuthWhiteList,
		serverutil.NewServerManager,
		serverutil.GetHTTPServer,
		serverutil.GetGRPCServer,
		serviceexport.ExportServices,
		serverutil.TODOAppServices,
	))
	return nil, func() {}, nil
}
