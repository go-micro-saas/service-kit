package configutil

import (
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
)

var (
	_bootstrap *Bootstrap
)

func SetBootstrap(bootstrap *Bootstrap) {
	_bootstrap = bootstrap
}

func GetBootstrap() (*Bootstrap, error) {
	if _bootstrap == nil {
		e := errorpkg.ErrorUninitialized("bootstrap is uninitialized")
		return nil, errorpkg.WithStack(e)
	}
	return _bootstrap, nil
}
