package redisutil

import (
	"github.com/redis/go-redis/v9"
	"sync"

	configpb "github.com/go-micro-saas/service-kit/api/config"
	"github.com/google/wire"
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

func GetRedisClient(redisManager RedisManager) (redis.UniversalClient, error) {
	return redisManager.GetClient()
}
