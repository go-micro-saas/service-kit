package postgresutil

import (
	configpb "github.com/go-micro-saas/service-kit/api/config"
	loggerutil "github.com/go-micro-saas/service-kit/setup/logger"
	"github.com/google/wire"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	"sync"
)

var ProviderSet = wire.NewSet(NewSingletonPostgresManager)

var (
	singletonMutex           sync.Once
	singletonPostgresManager PostgresManager
)

func NewSingletonPostgresManager(conf *configpb.PSQL, loggerManager loggerutil.LoggerManager) (PostgresManager, error) {
	var err error
	singletonMutex.Do(func() {
		singletonPostgresManager, err = NewPostgresManager(conf, loggerManager)
		if err != nil {
			singletonMutex = sync.Once{}
		}
	})
	return singletonPostgresManager, err
}

func GetPostgresManager() (PostgresManager, error) {
	if singletonPostgresManager == nil {
		e := errorpkg.ErrorUnimplemented("")
		return nil, errorpkg.WithStack(e)
	}
	return singletonPostgresManager, nil
}
