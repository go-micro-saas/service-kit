package dbutil

import setuputil "github.com/go-micro-saas/service-kit/setup"

type MigrationFunc func(launcherManager setuputil.LauncherManager, opts ...MigrationOptions)

// MigrationOptions ...
type MigrationOptions struct {
	Close bool
}

// MigrationOption ...
type MigrationOption func(*MigrationOptions)

// WithClose 退出后关闭资源
func WithClose() MigrationOption {
	return func(o *MigrationOptions) {
		o.Close = true
	}
}
