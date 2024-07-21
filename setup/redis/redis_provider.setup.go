package redisutil

import (
	configpb "github.com/go-micro-saas/service-kit/api/config"
	"github.com/google/wire"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	"sync"
)

var ProviderSet = wire.NewSet(NewSingletonRedisManager)

var (
	singletonMutex        sync.Once
	singletonRedisManager RedisManager
)

func NewSingletonRedisManager(conf *configpb.Redis) (RedisManager, error) {
	var err error
	singletonMutex.Do(func() {
		singletonRedisManager, err = NewRedisManager(conf)
		if err != nil {
			singletonMutex = sync.Once{}
		}
	})
	return singletonRedisManager, err
}

func GetRedisManager() (RedisManager, error) {
	if singletonRedisManager == nil {
		e := errorpkg.ErrorUnimplemented("")
		return nil, errorpkg.WithStack(e)
	}
	return singletonRedisManager, nil
}
