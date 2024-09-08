package dbutil

// MigrationOptions ...
type MigrationOptions struct {
	close bool
}

// MigrationOption ...
type MigrationOption func(*MigrationOptions)

// WithClose 退出后关闭资源
func WithClose() MigrationOption {
	return func(o *MigrationOptions) {
		o.close = true
	}
}
