// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package runservices

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-micro-saas/service-kit/config"
	"github.com/go-micro-saas/service-kit/server"
	"github.com/go-micro-saas/service-kit/setup"
	"github.com/go-micro-saas/service-kit/testdata/ping-service/api"
	"github.com/go-micro-saas/service-kit/testdata/ping-service/cmd/ping-service/export"
)

// Injectors from wire.go:

func initServiceApp(configFilePath string, configOpts ...configutil.Option) (*kratos.App, func(), error) {
	launcherManager, cleanup, err := setuputil.NewLauncherManagerWithCleanup(configFilePath, configOpts...)
	if err != nil {
		return nil, nil, err
	}
	v := api.GetAuthWhiteList()
	serverManager, err := serverutil.NewServerManager(launcherManager, v)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	server, err := serverutil.GetHTTPServer(serverManager)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	grpcServer, err := serverutil.GetGRPCServer(serverManager)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	services, cleanup2, err := exportservices.ExportServices(launcherManager, server, grpcServer)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	app, cleanup3, err := serverutil.TODOAppServices(serverManager, services)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return app, func() {
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}
