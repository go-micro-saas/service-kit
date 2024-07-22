package mysqlutil

import (
	"sync"

	configpb "github.com/go-micro-saas/service-kit/api/config"
	loggerutil "github.com/go-micro-saas/service-kit/setup/logger"
	"github.com/google/wire"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
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

func GetMysqlManager() (MysqlManager, error) {
	if singletonMysqlManager == nil {
		e := errorpkg.ErrorUninitialized("")
		return nil, errorpkg.WithStack(e)
	}
	return singletonMysqlManager, nil
}
