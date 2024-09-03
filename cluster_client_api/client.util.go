package clientutil

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	configpb "github.com/go-micro-saas/service-kit/api/config"
	setuputil "github.com/go-micro-saas/service-kit/setup"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	registrypkg "github.com/ikaiguang/go-srv-kit/kratos/registry"
	stdgrpc "google.golang.org/grpc"
	"sync"
)

type ClientAPIManager interface {
	// RegisterServiceAPIConfigs 注册服务API，覆盖更新
	RegisterServiceAPIConfigs(apis []*configpb.ClusterClientApi) error
	// GetServiceAPIConfig serviceName统一使用常量定义
	GetServiceAPIConfig(serviceName ServiceName) (*Config, error)
	// NewGRPCConnection grpc 链接
	NewGRPCConnection(logger log.Logger, serviceName ServiceName, otherOpts ...grpc.ClientOption) (*stdgrpc.ClientConn, error)
	// NewHTTPClient http 客户端
	NewHTTPClient(logger log.Logger, serviceName ServiceName, otherOpts ...http.ClientOption) (*http.Client, error)
}

type apiManager struct {
	launcherManager setuputil.LauncherManager
	configMap       map[ServiceName]*Config
	configMutex     sync.RWMutex
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
		consulClient, err := s.launcherManager.GetConsulClient()
		if err != nil {
			return nil, err
		}
		r, err := registrypkg.NewConsulRegistry(consulClient)
		if err != nil {
			return nil, err
		}
		return r, nil
	case configpb.ClusterClientApi_RT_ETCD:
		e := errorpkg.ErrorUnimplemented(apiConfig.RegistryType.String())
		return nil, errorpkg.WithStack(e)
	}
}
