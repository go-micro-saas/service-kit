package setuputil

import (
	"github.com/go-kratos/kratos/contrib/config/consul/v2"
	"github.com/go-kratos/kratos/v2/config"
	configpb "github.com/go-micro-saas/service-kit/api/config"
	apputil "github.com/go-micro-saas/service-kit/app"
	configutil "github.com/go-micro-saas/service-kit/config"
	consulutil "github.com/go-micro-saas/service-kit/setup/consul"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	stdlog "log"
	"strings"
)

const (
	CONFIG_METHOD_LOCAL  = "local"
	CONFIG_METHOD_CONSUL = "consul"
)

func LoadingConfig(filePath string) (*configpb.Bootstrap, error) {
	bootstrap, err := configutil.LoadingFile(filePath)
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
		consulManager, err := consulutil.NewConsulManager(bootstrap.GetConsul())
		if err != nil {
			return nil, err
		}
		return LoadingConfigFromConsul(consulManager, bootstrap.GetApp())
	}
}

func LoadingConfigFromConsul(consulManager consulutil.ConsulManager, appConfig *configpb.App) (*configpb.Bootstrap, error) {
	stdlog.Println("|==================== LOADING CONFIGURATION FROM: START ====================|")
	defer stdlog.Println()
	defer stdlog.Println("|==================== LOADING CONFIGURATION FROM: END ====================|")

	consulClient, err := consulManager.GetClient()
	if err != nil {
		return nil, err
	}

	// 配置source
	consulKeyPath := apputil.ConfigPath(appConfig)
	stdlog.Println("|*** LOADING: path to consul configuration: ", consulKeyPath)
	cs, err := consul.New(consulClient, consul.WithPath(consulKeyPath))
	if err != nil {
		e := errorpkg.ErrorInternalError(err.Error())
		return nil, errorpkg.WithStack(e)
	}

	var opts []config.Option
	stdlog.Println("|*** LOADING: consul source: ...")
	opts = append(opts, config.WithSource(cs))

	handler := config.New(opts...)
	defer func() {
		_ = handler.Close()
	}()

	// 加载配置
	if err = handler.Load(); err != nil {
		err = errorpkg.WithStack(errorpkg.ErrorInternalError(err.Error()))
		return nil, err
	}

	// 读取配置文件
	conf := &configpb.Bootstrap{}
	if err = handler.Scan(conf); err != nil {
		err = errorpkg.WithStack(errorpkg.ErrorInternalError(err.Error()))
		return nil, err
	}

	return conf, nil
}
