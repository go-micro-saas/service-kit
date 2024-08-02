package apputil

import (
	configpb "github.com/go-micro-saas/service-kit/api/config"
	apppkg "github.com/ikaiguang/go-srv-kit/kratos/app"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	"sync"
)

var (
	_bootstrap *configpb.Bootstrap

	// 不要直接使用 s.env, 请使用 Env()
	_env     apppkg.RuntimeEnvEnum_RuntimeEnv
	_envOnce sync.Once
)

func SetConfig(bootstrap *configpb.Bootstrap) {
	_bootstrap = bootstrap
}

func GetConfig() (*configpb.Bootstrap, error) {
	if _bootstrap == nil {
		e := errorpkg.ErrorUninitialized("bootstrap is uninitialized")
		return nil, errorpkg.WithStack(e)
	}
	return _bootstrap, nil
}

func Env() apppkg.RuntimeEnvEnum_RuntimeEnv {
	return apppkg.ParseEnv(_bootstrap.GetApp().GetServerEnv())
}

func IsDebugMode() bool {
	switch Env() {
	default:
		return false
	case apppkg.RuntimeEnvEnum_LOCAL, apppkg.RuntimeEnvEnum_DEVELOP, apppkg.RuntimeEnvEnum_TESTING:
		return true
	}
}

func IsLocalMode() bool {
	switch Env() {
	default:
		return false
	case apppkg.RuntimeEnvEnum_LOCAL:
		return true
	}
}
