package loggerutil

import (
	configpb "github.com/go-micro-saas/service-kit/api/config"
	configutil "github.com/go-micro-saas/service-kit/config"
	"github.com/google/wire"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	"sync"
)

var ProviderSet = wire.NewSet(NewSingletonLoggerManager)

var (
	singletonMutex         sync.Once
	singletonLoggerManager LoggerManager
)

func NewSingletonLoggerManager(conf *configpb.Log, appConfig *configpb.App) (LoggerManager, error) {
	var err error
	singletonMutex.Do(func() {
		singletonLoggerManager, err = NewLoggerManager(conf, appConfig)
		if err != nil {
			singletonMutex = sync.Once{}
		}
	})
	return singletonLoggerManager, err
}

func GetLoggerManager() (LoggerManager, error) {
	//if singletonLoggerManager == nil {
	//	return AttemptNewLoggerManager()
	//}
	if singletonLoggerManager == nil {
		e := errorpkg.ErrorUninitialized("")
		return nil, errorpkg.WithStack(e)
	}
	return singletonLoggerManager, nil
}

// AttemptNewLoggerManager 尝试初始化
func AttemptNewLoggerManager() (LoggerManager, error) {
	return NewSingletonLoggerManager(configutil.LogConfig(), configutil.AppConfig())
}
