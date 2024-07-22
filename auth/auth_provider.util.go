package authutil

import (
	"sync"

	configpb "github.com/go-micro-saas/service-kit/api/config"
	loggerutil "github.com/go-micro-saas/service-kit/logger"
	"github.com/google/wire"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	"github.com/redis/go-redis/v9"
)

var ProviderSet = wire.NewSet(NewSingletonAuthManager)

var (
	singletonMutex       sync.Once
	singletonAuthManager AuthManager
)

func NewSingletonAuthManager(
	conf *configpb.Encrypt_TokenEncrypt,
	redisCC redis.UniversalClient,
	loggerManager loggerutil.LoggerManager,
) (AuthManager, error) {
	var err error
	singletonMutex.Do(func() {
		singletonAuthManager, err = NewAuthManager(conf, redisCC, loggerManager)
		if err != nil {
			singletonMutex = sync.Once{}
		}
	})
	return singletonAuthManager, err
}

func GetAuthManager() (AuthManager, error) {
	if singletonAuthManager == nil {
		e := errorpkg.ErrorUninitialized("")
		return nil, errorpkg.WithStack(e)
	}
	return singletonAuthManager, nil
}
