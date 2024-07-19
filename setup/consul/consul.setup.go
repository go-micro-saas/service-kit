package consulutil

import (
	configpb "github.com/go-micro-saas/service-kit/api/config"
	consulapi "github.com/hashicorp/consul/api"
	consulpkg "github.com/ikaiguang/go-srv-kit/data/consul"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	stdlog "log"
	"sync"
)

type consulManager struct {
	conf *configpb.Consul

	consulOnce   sync.Once
	consulClient *consulapi.Client
}

type ConsulManager interface {
	GetClient() (*consulapi.Client, error)
}

func NewConsulManager(conf *configpb.Consul) (ConsulManager, error) {
	if conf == nil {
		e := errorpkg.ErrorBadRequest("[请配置服务再启动] config key : consul")
		return nil, errorpkg.WithStack(e)
	}
	return &consulManager{
		conf: conf,
	}, nil
}

func (s *consulManager) GetClient() (*consulapi.Client, error) {
	var err error
	s.consulOnce.Do(func() {
		s.consulClient, err = s.loadingConsulClient()
		if err != nil {
			s.consulOnce = sync.Once{}
		}
	})
	return s.consulClient, err
}

func (s *consulManager) loadingConsulClient() (*consulapi.Client, error) {
	stdlog.Println("|*** 加载：Consul客户端：...")
	cc, err := consulpkg.NewConsulClient(ToConsulConfig(s.conf))
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
