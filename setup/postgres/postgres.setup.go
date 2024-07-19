package postgresutil

import (
	configpb "github.com/go-micro-saas/service-kit/api/config"
	loggerutil "github.com/go-micro-saas/service-kit/setup/logger"
	gormpkg "github.com/ikaiguang/go-srv-kit/data/gorm"
	psqlpkg "github.com/ikaiguang/go-srv-kit/data/postgres"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sync"
)

type postgresManager struct {
	conf          *configpb.PSQL
	loggerManager loggerutil.LoggerManager

	postgresOnce   sync.Once
	postgresClient *gorm.DB
}

type PostgresManager interface {
	GetDB() (*gorm.DB, error)
}

func NewPostgresManager(conf *configpb.PSQL, loggerManager loggerutil.LoggerManager) (PostgresManager, error) {
	if conf == nil {
		e := errorpkg.ErrorBadRequest("[请配置服务再启动] config key : psql")
		return nil, errorpkg.WithStack(e)
	}
	return &postgresManager{
		conf:          conf,
		loggerManager: loggerManager,
	}, nil
}

func (s *postgresManager) GetDB() (*gorm.DB, error) {
	var err error
	s.postgresOnce.Do(func() {
		s.postgresClient, err = s.loadingPostgresDB()
		if err != nil {
			s.postgresOnce = sync.Once{}
		}
	})
	return s.postgresClient, err
}

func (s *postgresManager) loadingPostgresDB() (*gorm.DB, error) {
	// logger
	var (
		writers = make([]logger.Writer, 0, 2)
	)
	if s.loggerManager.EnableConsole() {
		writers = append(writers, gormpkg.NewStdWriter())
	}
	if s.loggerManager.EnableFile() {
		writer, err := s.loggerManager.GetWriter()
		if err != nil {
			return nil, err
		}
		writers = append(writers, gormpkg.NewJSONWriter(writer))
	}

	var opts = make([]gormpkg.Option, 0, 1)
	if len(writers) > 0 {
		opts = append(opts, gormpkg.WithWriters(writers...))
	}
	return psqlpkg.NewDB(ToPSQLConfig(s.conf), opts...)
}

// ToPSQLConfig ...
func ToPSQLConfig(cfg *configpb.PSQL) *psqlpkg.Config {
	return &psqlpkg.Config{
		Dsn:             cfg.Dsn,
		SlowThreshold:   cfg.SlowThreshold,
		LoggerEnable:    cfg.LoggerEnable,
		LoggerColorful:  cfg.LoggerColorful,
		LoggerLevel:     cfg.LoggerLevel,
		ConnMaxActive:   cfg.ConnMaxActive,
		ConnMaxLifetime: cfg.ConnMaxLifetime,
		ConnMaxIdle:     cfg.ConnMaxIdle,
		ConnMaxIdleTime: cfg.ConnMaxIdleTime,
	}
}
