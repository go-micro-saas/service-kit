package clientutil

import (
	"github.com/go-kratos/kratos/v2/registry"
	configpb "github.com/go-micro-saas/service-kit/api/config"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	logpkg "github.com/ikaiguang/go-srv-kit/kratos/log"
	registrypkg "github.com/ikaiguang/go-srv-kit/kratos/registry"
	"sync"
)

type clientAPIManager struct {
	opt         *option
	configMap   map[ServiceName]*Config
	configMutex sync.RWMutex
}

func NewClientAPIManager(opts ...Option) (ClientAPIManager, error) {
	o := &option{}
	o.logger, _ = logpkg.NewDummyLogger()
	for i := range opts {
		opts[i](o)
	}
	return &clientAPIManager{
		opt:         o,
		configMap:   nil,
		configMutex: sync.RWMutex{},
	}, nil
}

// RegisterServiceAPIConfigs 注册服务API，覆盖已有服务
func (s *clientAPIManager) RegisterServiceAPIConfigs(apiConfigs []*configpb.ClusterClientApi, opts ...Option) error {
	for i := range opts {
		opts[i](s.opt)
	}

	s.configMutex.Lock()
	defer s.configMutex.Unlock()

	var (
		hasConsulRegistry, hasEtcdRegistry bool
	)
	for i := range apiConfigs {
		if err := apiConfigs[i].Validate(); err != nil {
			e := errorpkg.ErrorBadRequest("")
			return errorpkg.Wrap(e, err)
		}
		conf := &Config{}
		conf.SetByPbClusterClientApi(apiConfigs[i])
		s.configMap[ServiceName(apiConfigs[i].ServiceName)] = conf
		if conf.IsConsulRegistry() {
			hasConsulRegistry = true
		} else if conf.IsEtcdRegistry() {
			hasEtcdRegistry = true
		}
	}
	if hasConsulRegistry && s.opt.consulClient == nil {
		return errorpkg.WithStack(uninitializedConsulClientError)
	}
	if hasEtcdRegistry && s.opt.etcdClient == nil {
		return errorpkg.WithStack(uninitializedEtcdClientError)
	}
	return nil
}

func (s *clientAPIManager) NewAPIConnection(serviceName ServiceName) (APIConnection, error) {
	apiConfig, err := s.GetServiceAPIConfig(serviceName)
	if err != nil {
		return nil, err
	}
	conn := &clientConnection{}
	conn.SetTransportType(apiConfig.TransportType)
	switch apiConfig.TransportType {
	default:
		conn.httpClient, err = s.NewHTTPClient(apiConfig)
		if err != nil {
			return nil, err
		}
	case configpb.ClusterClientApi_TT_HTTP:
		conn.httpClient, err = s.NewHTTPClient(apiConfig)
		if err != nil {
			return nil, err
		}
	case configpb.ClusterClientApi_TT_GRPC:
		conn.grpcConn, err = s.NewGRPCConnection(apiConfig)
		if err != nil {
			return nil, err
		}
	}
	return conn, nil
}

func (s *clientAPIManager) GetServiceAPIConfig(serviceName ServiceName) (*Config, error) {
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

func (s *clientAPIManager) getRegistryDiscovery(apiConfig *Config) (registry.Discovery, error) {
	switch apiConfig.RegistryType {
	default:
		e := errorpkg.ErrorUnimplemented(apiConfig.RegistryType.String())
		return nil, errorpkg.WithStack(e)
	case configpb.ClusterClientApi_RT_CONSUL:
		if s.opt.consulClient == nil {
			return nil, errorpkg.WithStack(uninitializedConsulClientError)
		}
		r, err := registrypkg.NewConsulRegistry(s.opt.consulClient)
		if err != nil {
			return nil, err
		}
		return r, nil
	case configpb.ClusterClientApi_RT_ETCD:
		if s.opt.etcdClient == nil {
			return nil, errorpkg.WithStack(uninitializedEtcdClientError)
		}
		r, err := registrypkg.NewEtcdRegistry(s.opt.etcdClient)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
}
