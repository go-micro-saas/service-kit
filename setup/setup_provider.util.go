package setuputil

import (
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/go-kratos/kratos/v2/log"
	consulapi "github.com/hashicorp/consul/api"
	authpkg "github.com/ikaiguang/go-srv-kit/kratos/auth"
	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"gorm.io/gorm"
	"sync"
)

var (
	singletonMutex           sync.Once
	singletonLauncherManager LauncherManager
)

func NewSingletonLauncherManager(configFilePath string) (LauncherManager, error) {
	var err error
	singletonMutex.Do(func() {
		singletonLauncherManager, err = NewLauncherManager(configFilePath)
	})
	if err != nil {
		singletonMutex = sync.Once{}
	}
	return singletonLauncherManager, err
}

func GetLogger(launcherManager LauncherManager) (log.Logger, error) {
	return launcherManager.GetLogger()
}
func GetLoggerForMiddleware(launcherManager LauncherManager) (log.Logger, error) {
	return launcherManager.GetLoggerForMiddleware()
}
func GetLoggerForHelper() (log.Logger, error) {
	return GetLoggerForHelper()
}
func GetRedisClient(launcherManager LauncherManager) (redis.UniversalClient, error) {
	return launcherManager.GetRedisClient()
}
func GetMysqlDBConn(launcherManager LauncherManager) (*gorm.DB, error) {
	return launcherManager.GetMysqlDBConn()
}
func GetPostgresDBConn(launcherManager LauncherManager) (*gorm.DB, error) {
	return launcherManager.GetPostgresDBConn()
}
func GetConsulClient(launcherManager LauncherManager) (*consulapi.Client, error) {
	return launcherManager.GetConsulClient()
}
func GetJaegerExporter(launcherManager LauncherManager) (*jaeger.Exporter, error) {
	return launcherManager.GetJaegerExporter()
}
func GetRabbitmqConn(launcherManager LauncherManager) (*amqp.ConnectionWrapper, error) {
	return launcherManager.GetRabbitmqConn()
}
func GetTokenManager(launcherManager LauncherManager) (authpkg.TokenManger, error) {
	return launcherManager.GetTokenManager()
}
func GetAuthManager(launcherManager LauncherManager) (authpkg.AuthRepo, error) {
	return launcherManager.GetAuthManager()
}
func Close(launcherManager LauncherManager) error {
	return launcherManager.Close()
}
