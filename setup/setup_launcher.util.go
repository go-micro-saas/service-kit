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
	stdlog "log"
	"sync"
)

type launcherManager struct {
	conf *configpb.Bootstrap

	loggerManagerOnce   sync.Once
	loggerManager       loggerutil.LoggerManager
	redisManagerOnce    sync.Once
	redisManager        redisutil.RedisManager
	mysqlManagerOnce    sync.Once
	mysqlManager        mysqlutil.MysqlManager
	postgresManagerOnce sync.Once
	postgresManager     postgresutil.PostgresManager
	consulManagerOnce   sync.Once
	consulManager       consulutil.ConsulManager
	jaegerManagerOnce   sync.Once
	jaegerManager       jaegerutil.JaegerManager
	rabbitmqManagerOnce sync.Once
	rabbitmqManager     rabbitmqutil.RabbitmqManager
	authInstanceOnce    sync.Once
	authInstance        authutil.AuthInstance
}

func (s *launcherManager) getLoggerManager() (loggerutil.LoggerManager, error) {
	logConfig := s.conf.GetLog()
	appConfig := s.conf.GetApp()
	loggerManager, err := loggerutil.NewSingletonLoggerManager(logConfig, appConfig)
	if err != nil {
		return nil, err
	}
	s.loggerManager = loggerManager
	return loggerManager, nil
}

func (s *launcherManager) getSingletonLoggerManager() (loggerutil.LoggerManager, error) {
	var err error
	s.loggerManagerOnce.Do(func() {
		s.loggerManager, err = s.getLoggerManager()
	})
	if err != nil {
		s.loggerManagerOnce = sync.Once{}
	}
	return s.loggerManager, err
}

func (s *launcherManager) GetLogger() (log.Logger, error) {
	loggerManager, err := s.getSingletonLoggerManager()
	if err != nil {
		return nil, err
	}
	return loggerManager.GetLogger()
}

func (s *launcherManager) GetLoggerForMiddleware() (log.Logger, error) {
	loggerManager, err := s.getSingletonLoggerManager()
	if err != nil {
		return nil, err
	}
	return loggerManager.GetLoggerForMiddleware()
}

func (s *launcherManager) GetLoggerForHelper() (log.Logger, error) {
	loggerManager, err := s.getSingletonLoggerManager()
	if err != nil {
		return nil, err
	}
	return loggerManager.GetLoggerForHelper()
}

func (s *launcherManager) getRedisManager() (redisutil.RedisManager, error) {
	redisConfig := s.conf.GetRedis()
	redisManager, err := redisutil.NewSingletonRedisManager(redisConfig)
	if err != nil {
		return nil, err
	}
	s.redisManager = redisManager
	return redisManager, nil
}

func (s *launcherManager) getSingletonRedisManager() (redisutil.RedisManager, error) {
	var err error
	s.redisManagerOnce.Do(func() {
		s.redisManager, err = s.getRedisManager()
	})
	if err != nil {
		s.redisManagerOnce = sync.Once{}
	}
	return s.redisManager, err
}

func (s *launcherManager) GetRedisClient() (redis.UniversalClient, error) {
	redisManager, err := s.getSingletonRedisManager()
	if err != nil {
		return nil, err
	}
	return redisManager.GetClient()
}

func (s *launcherManager) getMysqlManager() (mysqlutil.MysqlManager, error) {
	loggerManager, err := s.getSingletonLoggerManager()
	if err != nil {
		return nil, err
	}
	mysqlConfig := s.conf.GetMysql()
	mysqlManager, err := mysqlutil.NewSingletonMysqlManager(mysqlConfig, loggerManager)
	if err != nil {
		return nil, err
	}
	s.mysqlManager = mysqlManager
	return mysqlManager, nil
}

func (s *launcherManager) getSingletonMysqlManager() (mysqlutil.MysqlManager, error) {
	var err error
	s.mysqlManagerOnce.Do(func() {
		s.mysqlManager, err = s.getMysqlManager()
	})
	if err != nil {
		s.mysqlManagerOnce = sync.Once{}
	}
	return s.mysqlManager, err
}

func (s *launcherManager) GetMysqlDBConn() (*gorm.DB, error) {
	mysqlManager, err := s.getSingletonMysqlManager()
	if err != nil {
		return nil, err
	}
	return mysqlManager.GetDB()
}

func (s *launcherManager) getPostgresManager() (postgresutil.PostgresManager, error) {
	loggerManager, err := s.getSingletonLoggerManager()
	if err != nil {
		return nil, err
	}
	psqlConfig := s.conf.GetPsql()
	postgresManager, err := postgresutil.NewSingletonPostgresManager(psqlConfig, loggerManager)
	if err != nil {
		return nil, err
	}
	s.postgresManager = postgresManager
	return postgresManager, nil
}

func (s *launcherManager) getSingletonPostgresManager() (postgresutil.PostgresManager, error) {
	var err error
	s.postgresManagerOnce.Do(func() {
		s.postgresManager, err = s.getPostgresManager()
	})
	if err != nil {
		s.postgresManagerOnce = sync.Once{}
	}
	return s.postgresManager, err
}

func (s *launcherManager) GetPostgresDBConn() (*gorm.DB, error) {
	postgresManager, err := s.getSingletonPostgresManager()
	if err != nil {
		return nil, err
	}
	return postgresManager.GetDB()
}

func (s *launcherManager) getConsulManager() (consulutil.ConsulManager, error) {
	consulConfig := s.conf.GetConsul()
	consulManager, err := consulutil.NewSingletonConsulManager(consulConfig)
	if err != nil {
		return nil, err
	}
	s.consulManager = consulManager
	return consulManager, nil
}

func (s *launcherManager) getSingletonConsulManager() (consulutil.ConsulManager, error) {
	var err error
	s.consulManagerOnce.Do(func() {
		s.consulManager, err = s.getConsulManager()
	})
	if err != nil {
		s.consulManagerOnce = sync.Once{}
	}
	return s.consulManager, err
}

func (s *launcherManager) GetConsulClient() (*consulapi.Client, error) {
	consulManager, err := s.getSingletonConsulManager()
	if err != nil {
		return nil, err
	}
	return consulManager.GetClient()
}

func (s *launcherManager) getJaegerManager() (jaegerutil.JaegerManager, error) {
	jaegerConfig := s.conf.GetJaeger()
	jaegerManager, err := jaegerutil.NewSingletonJaegerManager(jaegerConfig)
	if err != nil {
		return nil, err
	}
	s.jaegerManager = jaegerManager
	return jaegerManager, nil
}

func (s *launcherManager) getSingletonJaegerManager() (jaegerutil.JaegerManager, error) {
	var err error
	s.jaegerManagerOnce.Do(func() {
		s.jaegerManager, err = s.getJaegerManager()
	})
	if err != nil {
		s.jaegerManagerOnce = sync.Once{}
	}
	return s.jaegerManager, err
}

func (s *launcherManager) GetJaegerExporter() (*jaeger.Exporter, error) {
	jaegerManager, err := s.getSingletonJaegerManager()
	if err != nil {
		return nil, err
	}
	return jaegerManager.GetExporter()
}

func (s *launcherManager) getRabbitmqManager() (rabbitmqutil.RabbitmqManager, error) {
	loggerManager, err := s.getSingletonLoggerManager()
	if err != nil {
		return nil, err
	}
	rabbitmqConfig := s.conf.GetRabbitmq()
	rabbitmqManager, err := rabbitmqutil.NewSingletonRabbitmqManager(rabbitmqConfig, loggerManager)
	if err != nil {
		return nil, err
	}
	s.rabbitmqManager = rabbitmqManager
	return rabbitmqManager, nil
}

func (s *launcherManager) getSingletonRabbitmqManager() (rabbitmqutil.RabbitmqManager, error) {
	var err error
	s.rabbitmqManagerOnce.Do(func() {
		s.rabbitmqManager, err = s.getRabbitmqManager()
	})
	if err != nil {
		s.rabbitmqManagerOnce = sync.Once{}
	}
	return s.rabbitmqManager, err
}

func (s *launcherManager) GetRabbitmqConn() (*amqp.ConnectionWrapper, error) {
	rabbitmqManager, err := s.getSingletonRabbitmqManager()
	if err != nil {
		return nil, err
	}
	return rabbitmqManager.GetClient()
}

func (s *launcherManager) getAuthInstance() (authutil.AuthInstance, error) {
	// logger
	loggerManager, err := s.getSingletonLoggerManager()
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

func (s *launcherManager) getSingletonAuthInstance() (authutil.AuthInstance, error) {
	var err error
	s.authInstanceOnce.Do(func() {
		s.authInstance, err = s.getAuthInstance()
	})
	if err != nil {
		s.authInstanceOnce = sync.Once{}
	}
	return s.authInstance, err
}

func (s *launcherManager) GetTokenManager() (authpkg.TokenManger, error) {
	authInstance, err := s.getSingletonAuthInstance()
	if err != nil {
		return nil, err
	}
	return authInstance.GetTokenManger()
}

func (s *launcherManager) GetAuthManager() (authpkg.AuthRepo, error) {
	authInstance, err := s.getSingletonAuthInstance()
	if err != nil {
		return nil, err
	}
	return authInstance.GetAuthManger()
}

func (s *launcherManager) Close() error {
	// 退出程序
	stdlog.Println("|==================== EXIT PROGRAM : START ====================|")
	defer stdlog.Println("|==================== EXIT PROGRAM : END ====================|")
	var errs []error

	// redis
	if s.redisManager != nil {
		if err := s.redisManager.Close(); err != nil {
			errs = append(errs, err)
		}
	}

	// mysql
	if s.mysqlManager != nil {
		if err := s.mysqlManager.Close(); err != nil {
			errs = append(errs, err)
		}
	}

	// postgres
	if s.postgresManager != nil {
		if err := s.postgresManager.Close(); err != nil {
			errs = append(errs, err)
		}
	}

	// consul
	if s.consulManager != nil {
		if err := s.consulManager.Close(); err != nil {
			errs = append(errs, err)
		}
	}

	// jaeger
	if s.jaegerManager != nil {
		if err := s.jaegerManager.Close(); err != nil {
			errs = append(errs, err)
		}
	}

	// rabbitmq
	if s.rabbitmqManager != nil {
		if err := s.rabbitmqManager.Close(); err != nil {
			errs = append(errs, err)
		}
	}

	// logger
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
