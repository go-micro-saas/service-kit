package setuputil

import (
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/go-kratos/kratos/v2/log"
	configpb "github.com/go-micro-saas/service-kit/api/config"
	apputil "github.com/go-micro-saas/service-kit/app"
	authutil "github.com/go-micro-saas/service-kit/auth"
	configutil "github.com/go-micro-saas/service-kit/config"
	consulutil "github.com/go-micro-saas/service-kit/consul"
	jaegerutil "github.com/go-micro-saas/service-kit/jaeger"
	loggerutil "github.com/go-micro-saas/service-kit/logger"
	mysqlutil "github.com/go-micro-saas/service-kit/mysql"
	postgresutil "github.com/go-micro-saas/service-kit/postgres"
	rabbitmqutil "github.com/go-micro-saas/service-kit/rabbitmq"
	redisutil "github.com/go-micro-saas/service-kit/redis"
	consulapi "github.com/hashicorp/consul/api"
	debugpkg "github.com/ikaiguang/go-srv-kit/debug"
	authpkg "github.com/ikaiguang/go-srv-kit/kratos/auth"
	logpkg "github.com/ikaiguang/go-srv-kit/kratos/log"
	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"gorm.io/gorm"
	stdlog "log"
)

type LauncherManager interface {
	GetConfig() *configpb.Bootstrap

	GetLogger() (log.Logger, error)
	GetLoggerForMiddleware() (log.Logger, error)
	GetLoggerForHelper() (log.Logger, error)

	GetRedisClient() (redis.UniversalClient, error)
	GetMysqlDBConn() (*gorm.DB, error)
	GetPostgresDBConn() (*gorm.DB, error)
	GetConsulClient() (*consulapi.Client, error)
	GetJaegerExporter() (*jaeger.Exporter, error)
	GetRabbitmqConn() (*amqp.ConnectionWrapper, error)

	GetTokenManager() (authpkg.TokenManger, error)
	GetAuthManager() (authpkg.AuthRepo, error)

	Close() error
}

func LoadingConfig(configFilePath string, configOpts ...configutil.Option) (*configpb.Bootstrap, error) {
	conf, err := configutil.Loading(configFilePath, configOpts...)
	if err != nil {
		return nil, err
	}
	apputil.SetConfig(conf)
	return conf, nil
}

func NewLauncherManagerWithCleanup(configFilePath string, configOpts ...configutil.Option) (LauncherManager, func(), error) {
	launcherManager, err := NewLauncherManager(configFilePath, configOpts...)
	if err != nil {
		return nil, nil, err
	}
	cleanup := func() {
		closeErr := launcherManager.Close()
		if closeErr != nil {
			stdlog.Printf("==> launcherManager.Close failed: %+v\n", closeErr)
		}
	}
	return launcherManager, cleanup, nil

}

func NewLauncherManager(configFilePath string, configOpts ...configutil.Option) (LauncherManager, error) {
	// 开始配置
	stdlog.Println("|==================== LOADING PROGRAM : START ====================|")
	defer stdlog.Println("|==================== LOADING PROGRAM : END ====================|")

	// 加载配置文件
	bootstrap, err := LoadingConfig(configFilePath, configOpts...)
	if err != nil {
		return nil, err
	}
	launcher := &launcherManager{
		conf: bootstrap,
	}

	// 初始化日志
	loggerManager, err := launcher.getSingletonLoggerManager()
	if err != nil {
		return nil, err
	}
	loggerForHelper, err := loggerManager.GetLoggerForHelper()
	if err != nil {
		return nil, err
	}
	logpkg.Setup(loggerForHelper)
	debugpkg.Setup(loggerForHelper)

	// redis
	redisConfig := bootstrap.GetRedis()
	if redisConfig.GetEnable() {
		_, err = launcher.GetRedisClient()
		if err != nil {
			return nil, err
		}
	}

	// mysql
	mysqlConfig := bootstrap.GetMysql()
	if mysqlConfig.GetEnable() {
		_, err = launcher.GetMysqlDBConn()
		if err != nil {
			return nil, err
		}
	}

	// postgres
	psqlConfig := bootstrap.GetPsql()
	if psqlConfig.GetEnable() {
		_, err = launcher.GetPostgresDBConn()
		if err != nil {
			return nil, err
		}
	}

	// consul
	consulConfig := bootstrap.GetConsul()
	if consulConfig.GetEnable() {
		_, err = launcher.GetConsulClient()
		if err != nil {
			return nil, err
		}
	}

	// jaeger
	jaegerConfig := bootstrap.GetJaeger()
	if jaegerConfig.GetEnable() {
		_, err = launcher.GetJaegerExporter()
		if err != nil {
			return nil, err
		}
	}

	// rabbitmq
	rabbitmqConfig := bootstrap.GetRabbitmq()
	if rabbitmqConfig.GetEnable() {
		_, err = launcher.GetRabbitmqConn()
		if err != nil {
			return nil, err
		}
	}

	// token
	settingConfig := bootstrap.GetSetting()
	if settingConfig.GetEnableAuthMiddleware() {
		_, err = launcher.getSingletonAuthInstance()
		if err != nil {
			return nil, err
		}
	}
	return launcher, nil
}

func testWireSetup(
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
