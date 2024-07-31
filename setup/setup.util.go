package setuputil

import (
	configpb "github.com/go-micro-saas/service-kit/api/config"
	authutil "github.com/go-micro-saas/service-kit/auth"
	configutil "github.com/go-micro-saas/service-kit/config"
	consulutil "github.com/go-micro-saas/service-kit/consul"
	jaegerutil "github.com/go-micro-saas/service-kit/jaeger"
	loggerutil "github.com/go-micro-saas/service-kit/logger"
	mysqlutil "github.com/go-micro-saas/service-kit/mysql"
	postgresutil "github.com/go-micro-saas/service-kit/postgres"
	rabbitmqutil "github.com/go-micro-saas/service-kit/rabbitmq"
	redisutil "github.com/go-micro-saas/service-kit/redis"
)

func LoadingConfig(configFilePath string) (*configpb.Bootstrap, error) {
	conf, err := configutil.Loading(configFilePath)
	if err != nil {
		return nil, err
	}
	configutil.SetConfig(conf)
	return conf, nil
}

type LauncherManager interface {
}

func Setup(
	conf *configpb.Bootstrap,
	loggerManager loggerutil.LoggerManager,
	mysqlManager mysqlutil.MysqlManager,
	postgresManager postgresutil.PostgresManager,
	redisManager redisutil.RedisManager,
	authInstance authutil.AuthInstance,
	consulManager consulutil.ConsulManager,
	jaegerManager jaegerutil.JaegerManager,
	rabbitmqManager rabbitmqutil.RabbitmqManager,
) (LauncherManager, error) {
	return nil, nil
}
