package jaegerutil

import (
	configpb "github.com/go-micro-saas/service-kit/api/config"
	"github.com/google/wire"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	"sync"
)

var ProviderSet = wire.NewSet(NewSingletonJaegerManager)

var (
	singletonMutex         sync.Once
	singletonJaegerManager JaegerManager
)

func NewSingletonJaegerManager(conf *configpb.Jaeger) (JaegerManager, error) {
	var err error
	singletonMutex.Do(func() {
		singletonJaegerManager, err = NewJaegerManager(conf)
		if err != nil {
			singletonMutex = sync.Once{}
		}
	})
	return singletonJaegerManager, err
}

func GetJaegerManager() (JaegerManager, error) {
	if singletonJaegerManager == nil {
		e := errorpkg.ErrorUnimplemented("")
		return nil, errorpkg.WithStack(e)
	}
	return singletonJaegerManager, nil
}
