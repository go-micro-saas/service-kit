package postgresutil

import (
	"gorm.io/gorm"
	"sync"

	configpb "github.com/go-micro-saas/service-kit/api/config"
	loggerutil "github.com/go-micro-saas/service-kit/logger"
	"github.com/google/wire"
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

func GetDBConn(postgresManager PostgresManager) (*gorm.DB, error) {
	return postgresManager.GetDB()
}
