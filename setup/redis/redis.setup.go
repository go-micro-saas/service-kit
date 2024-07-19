package redisutil

import (
	configpb "github.com/go-micro-saas/service-kit/api/config"
	loggerutil "github.com/go-micro-saas/service-kit/setup/logger"
	redispkg "github.com/ikaiguang/go-srv-kit/data/redis"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	"github.com/redis/go-redis/v9"
	"sync"
)

type redisManager struct {
	conf          *configpb.Redis
	loggerManager loggerutil.LoggerManager

	redisOnce   sync.Once
	redisClient redis.UniversalClient
}

type RedisManager interface {
	GetClient() (redis.UniversalClient, error)
}

func NewRedisManager(conf *configpb.Redis, loggerManager loggerutil.LoggerManager) (RedisManager, error) {
	if conf == nil {
		e := errorpkg.ErrorBadRequest("[请配置服务再启动] config key : redis")
		return nil, errorpkg.WithStack(e)
	}
	return &redisManager{
		conf:          conf,
		loggerManager: loggerManager,
	}, nil
}

func (s *redisManager) GetClient() (redis.UniversalClient, error) {
	var err error
	s.redisOnce.Do(func() {
		s.redisClient, err = s.loadingRedisClient()
		if err != nil {
			s.redisOnce = sync.Once{}
		}
	})
	return s.redisClient, err
}

func (s *redisManager) loadingRedisClient() (redis.UniversalClient, error) {
	uc, err := redispkg.NewDB(ToRedisConfig(s.conf))
	if err != nil {
		e := errorpkg.ErrorInternalError(err.Error())
		return nil, errorpkg.WithStack(e)
	}
	return uc, nil
}

// ToRedisConfig ...
func ToRedisConfig(cfg *configpb.Redis) *redispkg.Config {
	return &redispkg.Config{
		Addresses:       cfg.Addresses,
		Username:        cfg.Username,
		Password:        cfg.Password,
		Db:              cfg.Db,
		DialTimeout:     cfg.DialTimeout,
		ReadTimeout:     cfg.ReadTimeout,
		WriteTimeout:    cfg.WriteTimeout,
		ConnMaxActive:   cfg.ConnMaxActive,
		ConnMaxLifetime: cfg.ConnMaxLifetime,
		ConnMaxIdle:     cfg.ConnMaxIdle,
		ConnMinIdle:     cfg.ConnMinIdle,
		ConnMaxIdleTime: cfg.ConnMaxIdleTime,
	}
}
