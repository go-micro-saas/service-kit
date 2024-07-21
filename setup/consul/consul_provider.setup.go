package consulutil

import (
	configpb "github.com/go-micro-saas/service-kit/api/config"
	"github.com/google/wire"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	"sync"
)

var ProviderSet = wire.NewSet(NewSingletonConsulManager)

var (
	singletonMutex         sync.Once
	singletonConsulManager ConsulManager
)

func NewSingletonConsulManager(conf *configpb.Consul) (ConsulManager, error) {
	var err error
	singletonMutex.Do(func() {
		singletonConsulManager, err = NewConsulManager(conf)
		if err != nil {
			singletonMutex = sync.Once{}
		}
	})
	return singletonConsulManager, err
}

func GetConsulManager() (ConsulManager, error) {
	if singletonConsulManager == nil {
		e := errorpkg.ErrorUnimplemented("")
		return nil, errorpkg.WithStack(e)
	}
	return singletonConsulManager, nil
}
