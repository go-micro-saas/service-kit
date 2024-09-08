package dbutil

import (
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type MigrationFunc func(dbConn *gorm.DB, opts ...MigrationOptions)

// MigrationOptions ...
type MigrationOptions struct {
	Logger log.Logger
	Close  bool
}

func DefaultMigrationOptions() *MigrationOptions {
	return &MigrationOptions{
		Logger: log.DefaultLogger,
	}
}

// MigrationOption ...
type MigrationOption func(*MigrationOptions)

func WithLogger(logger log.Logger) MigrationOption {
	return func(o *MigrationOptions) {
		o.Logger = logger
	}
}

// WithClose 退出后关闭资源
func WithClose() MigrationOption {
	return func(o *MigrationOptions) {
		o.Close = true
	}
}
