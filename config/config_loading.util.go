package configutil

import (
	"path/filepath"
	"runtime"
	"strings"

	configpb "github.com/go-micro-saas/service-kit/api/config"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
)

const (
	CONFIG_METHOD_LOCAL  = "local"
	CONFIG_METHOD_CONSUL = "consul"
)

func Loading(filePath string) (*configpb.Bootstrap, error) {
	bootstrap, err := LoadingFile(filePath)
	if err != nil {
		return nil, err
	}
	if bootstrap.GetApp() == nil {
		e := errorpkg.ErrorBadRequest("[CONFIGURATION] config error, key = app")
		return nil, errorpkg.WithStack(e)
	}
	if bootstrap.GetConsul() == nil {
		e := errorpkg.ErrorBadRequest("[CONFIGURATION] config error, key = consul")
		return nil, errorpkg.WithStack(e)
	}
	method := strings.ToLower(bootstrap.GetApp().GetConfigMethod())
	switch method {
	default:
		return bootstrap, err
	case CONFIG_METHOD_CONSUL:
		//从consul加载配置
		consulClient, err := newConsulClient(bootstrap.GetConsul())
		if err != nil {
			return nil, err
		}
		return LoadingConfigFromConsul(consulClient, bootstrap.GetApp())
	}
}

func CurrentPath() string {
	_, file, _, _ := runtime.Caller(0)

	return filepath.Dir(file)
}
