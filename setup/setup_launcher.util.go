package setuputil

import (
	stderrors "errors"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/go-kratos/kratos/v2/log"
	configpb "github.com/go-micro-saas/service-kit/api/config"
	authutil "github.com/go-micro-saas/service-kit/auth"
	consulutil "github.com/go-micro-saas/service-kit/consul"
	jaegerutil "github.com/go-micro-saas/service-kit/jaeger"
	loggerutil "github.com/go-micro-saas/service-kit/logger"
	mysqlutil "github.com/go-micro-saas/service-kit/mysql"
	postgresutil "github.com/go-micro-saas/service-kit/postgres"
	rabbitmqutil "github.com/go-micro-saas/service-kit/rabbitmq"
	redisutil "github.com/go-micro-saas/service-kit/redis"
	consulapi "github.com/hashicorp/consul/api"
	authpkg "github.com/ikaiguang/go-srv-kit/kratos/auth"
	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"gorm.io/gorm"
)

type launcherManager struct {
	conf *configpb.Bootstrap

	loggerManager   loggerutil.LoggerManager
	redisManager    redisutil.RedisManager
	mysqlManager    mysqlutil.MysqlManager
	postgresManager postgresutil.PostgresManager
	consulManager   consulutil.ConsulManager
	jaegerManager   jaegerutil.JaegerManager
	rabbitmqManager rabbitmqutil.RabbitmqManager
	authInstance    authutil.AuthInstance
}

func (s *launcherManager) getLoggerManager() (loggerutil.LoggerManager, error) {
	logConfig := s.conf.GetLog()
	appConfig := s.conf.GetApp()
	var err error
	s.loggerManager, err = loggerutil.NewSingletonLoggerManager(logConfig, appConfig)
	if err != nil {
		return nil, err
	}
	return s.loggerManager, nil
}

func (s *launcherManager) GetLogger() (log.Logger, error) {
	loggerManager, err := s.getLoggerManager()
	if err != nil {
		return nil, err
	}
	return loggerManager.GetLogger()
}

func (s *launcherManager) GetLoggerForMiddleware() (log.Logger, error) {
	loggerManager, err := s.getLoggerManager()
	if err != nil {
		return nil, err
	}
	return loggerManager.GetLoggerForMiddleware()
}

func (s *launcherManager) GetLoggerForHelper() (log.Logger, error) {
	loggerManager, err := s.getLoggerManager()
	if err != nil {
		return nil, err
	}
	return loggerManager.GetLoggerForHelper()
}

func (s *launcherManager) GetRedisClient() (redis.UniversalClient, error) {
	redisConfig := s.conf.GetRedis()
	redisManager, err := redisutil.NewSingletonRedisManager(redisConfig)
	if err != nil {
		return nil, err
	}
	s.redisManager = redisManager
	return redisManager.GetClient()
}

func (s *launcherManager) GetMysqlDBConn() (*gorm.DB, error) {
	loggerManager, err := s.getLoggerManager()
	if err != nil {
		return nil, err
	}
	mysqlConfig := s.conf.GetMysql()
	mysqlManager, err := mysqlutil.NewSingletonMysqlManager(mysqlConfig, loggerManager)
	if err != nil {
		return nil, err
	}
	s.mysqlManager = mysqlManager
	return mysqlManager.GetDB()
}

func (s *launcherManager) GetPostgresDBConn() (*gorm.DB, error) {
	loggerManager, err := s.getLoggerManager()
	if err != nil {
		return nil, err
	}
	psqlConfig := s.conf.GetPsql()
	postgresManager, err := postgresutil.NewSingletonPostgresManager(psqlConfig, loggerManager)
	if err != nil {
		return nil, err
	}
	s.postgresManager = postgresManager
	return postgresManager.GetDB()
}

func (s *launcherManager) GetConsulClient() (*consulapi.Client, error) {
	consulConfig := s.conf.GetConsul()
	consulManager, err := consulutil.NewSingletonConsulManager(consulConfig)
	if err != nil {
		return nil, err
	}
	s.consulManager = consulManager
	return consulManager.GetClient()
}

func (s *launcherManager) GetJaegerExporter() (*jaeger.Exporter, error) {
	jaegerConfig := s.conf.GetJaeger()
	jaegerManager, err := jaegerutil.NewSingletonJaegerManager(jaegerConfig)
	if err != nil {
		return nil, err
	}
	s.jaegerManager = jaegerManager
	return jaegerManager.GetExporter()
}

func (s *launcherManager) GetRabbitmqConn() (*amqp.ConnectionWrapper, error) {
	loggerManager, err := s.getLoggerManager()
	if err != nil {
		return nil, err
	}
	rabbitmqConfig := s.conf.GetRabbitmq()
	rabbitmqManager, err := rabbitmqutil.NewSingletonRabbitmqManager(rabbitmqConfig, loggerManager)
	if err != nil {
		return nil, err
	}
	s.rabbitmqManager = rabbitmqManager
	return rabbitmqManager.GetClient()
}

func (s *launcherManager) getAuthInstance() (authutil.AuthInstance, error) {
	// logger
	loggerManager, err := s.getLoggerManager()
	if err != nil {
		return nil, err
	}
	// redis
	universalClient, err := s.GetRedisClient()
	if err != nil {
		return nil, err
	}
	// auth
	encryptTokenEncrypt := s.conf.GetEncrypt().GetTokenEncrypt()
	authInstance, err := authutil.NewSingletonAuthInstance(encryptTokenEncrypt, universalClient, loggerManager)
	if err != nil {
		return nil, err
	}
	s.authInstance = authInstance
	return authInstance, err
}

func (s *launcherManager) GetTokenManager() (authpkg.TokenManger, error) {
	authInstance, err := s.getAuthInstance()
	if err != nil {
		return nil, err
	}
	return authInstance.GetTokenManger()
}

func (s *launcherManager) GetAuthManager() (authpkg.AuthRepo, error) {
	authInstance, err := s.getAuthInstance()
	if err != nil {
		return nil, err
	}
	return authInstance.GetAuthManger()
}

func (s *launcherManager) Close() error {
	var errs []error
	if s.loggerManager != nil {
		if err := s.loggerManager.Close(); err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return stderrors.Join(errs...)
	}
	return nil
}
