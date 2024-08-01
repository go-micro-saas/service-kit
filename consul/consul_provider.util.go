package consulutil

import (
	consulapi "github.com/hashicorp/consul/api"
	"sync"

	configpb "github.com/go-micro-saas/service-kit/api/config"
)

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

func GetConsulClient(consulManager ConsulManager) (*consulapi.Client, error) {
	return consulManager.GetClient()
}
