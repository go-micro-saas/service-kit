package mysqlutil

import (
	configpb "github.com/go-micro-saas/service-kit/api/config"
	loggerutil "github.com/go-micro-saas/service-kit/setup/logger"
	gormpkg "github.com/ikaiguang/go-srv-kit/data/gorm"
	mysqlpkg "github.com/ikaiguang/go-srv-kit/data/mysql"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	stdlog "log"
	"sync"
)

type mysqlManager struct {
	conf          *configpb.MySQL
	loggerManager loggerutil.LoggerManager

	mysqlOnce   sync.Once
	mysqlClient *gorm.DB
}

type MysqlManager interface {
	GetDB() (*gorm.DB, error)
}

func NewMysqlManager(conf *configpb.MySQL, loggerManager loggerutil.LoggerManager) (MysqlManager, error) {
	if conf == nil {
		e := errorpkg.ErrorBadRequest("[请配置服务再启动] config key : mysql")
		return nil, errorpkg.WithStack(e)
	}
	return &mysqlManager{
		conf:          conf,
		loggerManager: loggerManager,
	}, nil
}

func (s *mysqlManager) GetDB() (*gorm.DB, error) {
	var err error
	s.mysqlOnce.Do(func() {
		s.mysqlClient, err = s.loadingMysqlDB()
		if err != nil {
			s.mysqlOnce = sync.Once{}
		}
	})
	return s.mysqlClient, err
}

func (s *mysqlManager) loadingMysqlDB() (*gorm.DB, error) {
	stdlog.Println("|*** 加载：MysqlDB：...")
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
	db, err := mysqlpkg.NewMysqlDB(ToMysqlConfig(s.conf), opts...)
	if err != nil {
		e := errorpkg.ErrorInternalError(err.Error())
		return nil, errorpkg.WithStack(e)
	}
	return db, nil
}

// ToMysqlConfig ...
func ToMysqlConfig(cfg *configpb.MySQL) *mysqlpkg.Config {
	return &mysqlpkg.Config{
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
