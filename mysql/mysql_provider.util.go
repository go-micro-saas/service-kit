package mysqlutil

import (
	"gorm.io/gorm"
	"sync"

	configpb "github.com/go-micro-saas/service-kit/api/config"
	loggerutil "github.com/go-micro-saas/service-kit/logger"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewSingletonMysqlManager)

var (
	singletonMutex        sync.Once
	singletonMysqlManager MysqlManager
)

func NewSingletonMysqlManager(conf *configpb.MySQL, loggerManager loggerutil.LoggerManager) (MysqlManager, error) {
	var err error
	singletonMutex.Do(func() {
		singletonMysqlManager, err = NewMysqlManager(conf, loggerManager)
		if err != nil {
			singletonMutex = sync.Once{}
		}
	})
	return singletonMysqlManager, err
}

func GetMysqlManager(mysqlManager MysqlManager) (*gorm.DB, error) {
	return mysqlManager.GetDB()
}
