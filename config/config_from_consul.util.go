package configutil

import (
	stdlog "log"

	"github.com/go-kratos/kratos/contrib/config/consul/v2"
	"github.com/go-kratos/kratos/v2/config"
	configpb "github.com/go-micro-saas/service-kit/api/config"
	apputil "github.com/go-micro-saas/service-kit/app"
	consulapi "github.com/hashicorp/consul/api"
	consulpkg "github.com/ikaiguang/go-srv-kit/data/consul"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
)

// LoadingConfigFromConsul 从consul中加载配置
// 首先：读取服务base配置
// 然后：读取本服务配置
// 最后：使用本服务配置 覆盖 base 配置
func LoadingConfigFromConsul(consulClient *consulapi.Client, appConfig *configpb.App) (*configpb.Bootstrap, error) {
	stdlog.Println("|==================== LOADING CONFIGURATION FROM: START ====================|")
	defer stdlog.Println()
	defer stdlog.Println("|==================== LOADING CONFIGURATION FROM: END ====================|")

	// 配置source
	consulKeyPath := apputil.Path(appConfig)
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
