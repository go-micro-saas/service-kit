package configutil

import (
	configs "github.com/go-micro-saas/service-kit/api/config"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
)

var (
	_bootstrap *configs.Bootstrap
)

func SetBootstrap(bootstrap *configs.Bootstrap) {
	_bootstrap = bootstrap
}

func GetBootstrap() (*configs.Bootstrap, error) {
	if _bootstrap == nil {
		e := errorpkg.ErrorUninitialized("bootstrap is nil")
		return nil, errorpkg.WrapWithEnumNumber(e, errorpkg.ERROR_UNINITIALIZED)
	}
	return _bootstrap, nil
}
