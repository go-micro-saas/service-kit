//go:build wireinject
// +build wireinject

package setuputil

import (
	configutil "github.com/go-micro-saas/service-kit/config"
	loggerutil "github.com/go-micro-saas/service-kit/logger"
	mysqlutil "github.com/go-micro-saas/service-kit/mysql"
	"github.com/google/wire"
)

var (
	LoggerProviderSet = wire.NewSet(configutil.LogConfig, configutil.AppConfig, loggerutil.NewSingletonLoggerManager)
	MysqlProviderSet  = wire.NewSet(configutil.MysqlConfig, mysqlutil.NewSingletonMysqlManager)
)

func setupLauncherManager(configFilePath string) (LauncherManager, error) {
	panic(wire.Build(LoadingConfig, LoggerProviderSet, MysqlProviderSet))
	return nil, nil
}
