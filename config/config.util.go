package configutil

import (
	configpb "github.com/go-micro-saas/service-kit/api/config"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
)

var (
	_bootstrap *configpb.Bootstrap
)

func SetBootstrap(bootstrap *configpb.Bootstrap) {
	_bootstrap = bootstrap
}

func GetBootstrap() (*configpb.Bootstrap, error) {
	if _bootstrap == nil {
		e := errorpkg.ErrorUninitialized("bootstrap is uninitialized")
		return nil, errorpkg.WithStack(e)
	}
	return _bootstrap, nil
}
