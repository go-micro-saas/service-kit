package rabbitmqutil

import (
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"sync"

	configpb "github.com/go-micro-saas/service-kit/api/config"
	loggerutil "github.com/go-micro-saas/service-kit/logger"
)

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

func GetRabbitmqConn(rabbitmqManager RabbitmqManager) (*amqp.ConnectionWrapper, error) {
	return rabbitmqManager.GetClient()
}
