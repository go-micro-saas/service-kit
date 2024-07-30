//go:build wireinject
// +build wireinject

package setuputil

import (
	"github.com/go-kratos/kratos/v2"
	configutil "github.com/go-micro-saas/service-kit/config"
	loggerutil "github.com/go-micro-saas/service-kit/logger"
	"github.com/google/wire"
)

// setupApp setup application.
func setupApp() (*kratos.App, func(), error) {
	panic(wire.Build(configutil.LogConfig, configutil.AppConfig, loggerutil.NewSingletonLoggerManager))
}
