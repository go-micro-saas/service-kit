package authutil

import (
	authpkg "github.com/ikaiguang/go-srv-kit/kratos/auth"
	"sync"

	configpb "github.com/go-micro-saas/service-kit/api/config"
	loggerutil "github.com/go-micro-saas/service-kit/logger"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
)

var ProviderSet = wire.NewSet(NewSingletonAuthInstance)

var (
	singletonMutex        sync.Once
	singletonAuthInstance AuthInstance
)

func NewSingletonAuthInstance(
	conf *configpb.Encrypt_TokenEncrypt,
	redisCC redis.UniversalClient,
	loggerManager loggerutil.LoggerManager,
) (AuthInstance, error) {
	var err error
	singletonMutex.Do(func() {
		singletonAuthInstance, err = NewAuthInstance(conf, redisCC, loggerManager)
		if err != nil {
			singletonMutex = sync.Once{}
		}
	})
	return singletonAuthInstance, err
}

func GetAuthManager(authInstance AuthInstance) (authpkg.AuthRepo, error) {
	return authInstance.GetAuthManger()
}

func GetTokenManger(authInstance AuthInstance) (authpkg.TokenManger, error) {
	return authInstance.GetTokenManger()
}
