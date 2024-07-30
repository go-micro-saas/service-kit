package configutil

import (
	stdlog "log"

	"github.com/go-kratos/kratos/contrib/config/consul/v2"
	"github.com/go-kratos/kratos/v2/config"
	configpb "github.com/go-micro-saas/service-kit/api/config"
	consulapi "github.com/hashicorp/consul/api"
	consulpkg "github.com/ikaiguang/go-srv-kit/data/consul"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
)

// LoadingConfigFromConsul 从consul中加载配置
// 首先：读取服务base配置
// 然后：读取本服务配置
// 最后：使用本服务配置 覆盖 base 配置
func LoadingConfigFromConsul(consulClient *consulapi.Client, appConfig *configpb.App) (*configpb.Bootstrap, error) {
	var bootstrap = &configpb.Bootstrap{}

	// 通用配置
	generalPath := appConfig.GetConfigPathForGeneral()
	if generalPath != "" {
		conf, err := loadingConfigFromConsul(consulClient, generalPath)
		if err != nil {
			return nil, err
		}
		bootstrap = conf
	} else {
		stdlog.Println("|*** INFO: no general configuration path configured")
	}

	// 服务配置 合并与覆盖
	serverPath := appConfig.GetConfigPathForServer()
	if serverPath != "" {
		conf, err := loadingConfigFromConsul(consulClient, serverPath)
		if err != nil {
			return nil, err
		}
		MergeConfig(bootstrap, conf)
	} else {
		stdlog.Println("|*** INFO: this service configuration path is not configured")
	}

	return bootstrap, nil
}

func loadingConfigFromConsul(consulClient *consulapi.Client, consulConfigPath string) (*configpb.Bootstrap, error) {
	stdlog.Println("|==================== LOADING CONSUL CONFIGURATION : START ====================|")
	defer stdlog.Println()
	defer stdlog.Println("|==================== LOADING CONSUL CONFIGURATION : END ====================|")
	stdlog.Println("|*** LOADING: consul configuration path: ", consulConfigPath)

	// 配置source
	cs, err := consul.New(consulClient, consul.WithPath(consulConfigPath))
	if err != nil {
		e := errorpkg.ErrorInternalError(err.Error())
		return nil, errorpkg.WithStack(e)
	}
	var opts []config.Option
	opts = append(opts, config.WithSource(cs))

	handler := config.New(opts...)
	defer func() {
		stdlog.Println("|*** LOADING: COMPLETE : consul configuration path: ", consulConfigPath)
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

func newConsulClient(cfg *configpb.Consul) (*consulapi.Client, error) {
	cc, err := consulpkg.NewConsulClient(ToConsulConfig(cfg))
	if err != nil {
		e := errorpkg.ErrorInternalError(err.Error())
		return nil, errorpkg.WithStack(e)
	}
	return cc, nil
}

// ToConsulConfig ...
func ToConsulConfig(cfg *configpb.Consul) *consulpkg.Config {
	return &consulpkg.Config{
		Scheme:             cfg.Scheme,
		Address:            cfg.Address,
		PathPrefix:         cfg.PathPrefix,
		Datacenter:         cfg.Datacenter,
		WaitTime:           cfg.WaitTime,
		Token:              cfg.Token,
		Namespace:          cfg.Namespace,
		Partition:          cfg.Partition,
		WithHttpBasicAuth:  cfg.WithHttpBasicAuth,
		AuthUsername:       cfg.AuthUsername,
		AuthPassword:       cfg.AuthPassword,
		InsecureSkipVerify: cfg.InsecureSkipVerify,
		TlsAddress:         cfg.TlsAddress,
		TlsCaPem:           cfg.TlsCaPem,
		TlsCertPem:         cfg.TlsCertPem,
		TlsKeyPem:          cfg.TlsKeyPem,
	}
}
