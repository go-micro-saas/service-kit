//go:build wireinject
// +build wireinject

package setuputil

import (
	authutil "github.com/go-micro-saas/service-kit/auth"
	configutil "github.com/go-micro-saas/service-kit/config"
	consulutil "github.com/go-micro-saas/service-kit/consul"
	jaegerutil "github.com/go-micro-saas/service-kit/jaeger"
	loggerutil "github.com/go-micro-saas/service-kit/logger"
	mysqlutil "github.com/go-micro-saas/service-kit/mysql"
	postgresutil "github.com/go-micro-saas/service-kit/postgres"
	rabbitmqutil "github.com/go-micro-saas/service-kit/rabbitmq"
	redisutil "github.com/go-micro-saas/service-kit/redis"
	"github.com/google/wire"
)

var (
	LoggerProviderSet   = wire.NewSet(configutil.LogConfig, configutil.AppConfig, loggerutil.NewSingletonLoggerManager)
	MysqlProviderSet    = wire.NewSet(configutil.MysqlConfig, mysqlutil.NewSingletonMysqlManager)
	PostgresProviderSet = wire.NewSet(configutil.PostgresConfig, postgresutil.NewSingletonPostgresManager)
	RedisProviderSet    = wire.NewSet(configutil.RedisConfig, redisutil.NewSingletonRedisManager)
	AuthProviderSet     = wire.NewSet(configutil.TokenEncryptConfig, redisutil.GetRedisClient, authutil.NewSingletonAuthInstance)
	ConsulProviderSet   = wire.NewSet(configutil.ConsulConfig, consulutil.NewSingletonConsulManager)
	JaegerProviderSet   = wire.NewSet(configutil.JaegerConfig, jaegerutil.NewSingletonJaegerManager)
	RabbitmqProviderSet = wire.NewSet(configutil.RabbitmqConfig, rabbitmqutil.NewSingletonRabbitmqManager)
)

func testWire(configFilePath string) (LauncherManager, error) {
	panic(wire.Build(
		LoadingConfig, LoggerProviderSet, MysqlProviderSet, PostgresProviderSet, RedisProviderSet,
		AuthProviderSet, ConsulProviderSet, JaegerProviderSet, RabbitmqProviderSet,
		testWireSetup,
	))
	return nil, nil
}
