package clientutil

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	configpb "github.com/go-micro-saas/service-kit/api/config"
	consulapi "github.com/hashicorp/consul/api"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	registrypkg "github.com/ikaiguang/go-srv-kit/kratos/registry"
	"sync"
)

type Option func(*option)

func WithConsulClient(consulClient *consulapi.Client) Option {
	return func(opt *option) {
		opt.consulClient = consulClient
	}
}

type option struct {
	consulClient *consulapi.Client
}

func NewClientAPIManager(logger log.Logger, opts ...Option) (ClientAPIManager, error) {
	o := &option{}
	for i := range opts {
		opts[i](o)
	}
	return &apiManager{
		logger:       logger,
		consulClient: o.consulClient,
		configMap:    nil,
		configMutex:  sync.RWMutex{},
	}, nil
}

type apiManager struct {
	logger       log.Logger
	consulClient *consulapi.Client
	configMap    map[ServiceName]*Config
	configMutex  sync.RWMutex
}

// RegisterServiceAPIConfigs 注册服务API，覆盖已有服务
func (s *apiManager) RegisterServiceAPIConfigs(apiConfigs []*configpb.ClusterClientApi) error {
	s.configMutex.Lock()
	defer s.configMutex.Unlock()
	for i := range apiConfigs {
		if err := apiConfigs[i].Validate(); err != nil {
			e := errorpkg.ErrorBadRequest("")
			return errorpkg.Wrap(e, err)
		}
		conf := &Config{}
		conf.SetByPbClusterClientApi(apiConfigs[i])
		s.configMap[ServiceName(apiConfigs[i].ServiceName)] = conf
	}
	return nil
}

func (s *apiManager) GetServiceAPIConfig(serviceName ServiceName) (*Config, error) {
	if serviceName.String() == "" {
		e := errorpkg.ErrorBadRequest("service name cannot be empty")
		return nil, errorpkg.WithStack(e)
	}
	s.configMutex.RLock()
	defer s.configMutex.RUnlock()
	conf, ok := s.configMap[serviceName]
	if !ok {
		e := errorpkg.ErrorRecordNotFound("service configuration not found; ServiceName: %s", serviceName.String())
		return nil, errorpkg.WithStack(e)
	}
	if conf == nil {
		e := errorpkg.ErrorInternalError("service configuration error: config == nil")
		return nil, errorpkg.WithStack(e)
	}
	return conf, nil
}

func (s *apiManager) getRegistryDiscovery(apiConfig *Config) (registry.Discovery, error) {
	switch apiConfig.RegistryType {
	default:
		e := errorpkg.ErrorUnimplemented(apiConfig.RegistryType.String())
		return nil, errorpkg.WithStack(e)
	case configpb.ClusterClientApi_RT_CONSUL:
		if s.consulClient == nil {
			e := errorpkg.ErrorBadRequest("consulClient == nil")
			return nil, errorpkg.WithStack(e)
		}
		r, err := registrypkg.NewConsulRegistry(s.consulClient)
		if err != nil {
			return nil, err
		}
		return r, nil
	case configpb.ClusterClientApi_RT_ETCD:
		e := errorpkg.ErrorUnimplemented(apiConfig.RegistryType.String())
		return nil, errorpkg.WithStack(e)
	}
}
