package rabbitmqutil

import (
	configpb "github.com/go-micro-saas/service-kit/api/config"
	loggerutil "github.com/go-micro-saas/service-kit/setup/logger"
	"github.com/google/wire"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	"sync"
)

var ProviderSet = wire.NewSet(NewSingletonRabbitmqManager)

var (
	singletonMutex           sync.Once
	singletonRabbitmqManager RabbitmqManager
)

func NewSingletonRabbitmqManager(conf *configpb.Rabbitmq, loggerManager loggerutil.LoggerManager) (RabbitmqManager, error) {
	var err error
	singletonMutex.Do(func() {
		singletonRabbitmqManager, err = NewRabbitmqManager(conf, loggerManager)
		if err != nil {
			singletonMutex = sync.Once{}
		}
	})
	return singletonRabbitmqManager, err
}

func GetRabbitmqManager() (RabbitmqManager, error) {
	if singletonRabbitmqManager == nil {
		e := errorpkg.ErrorUnimplemented("")
		return nil, errorpkg.WithStack(e)
	}
	return singletonRabbitmqManager, nil
}
