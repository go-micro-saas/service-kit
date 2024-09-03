package clientutil

import (
	configpb "github.com/go-micro-saas/service-kit/api/config"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	"sync"
)

type ClientAPIManager interface {
	// RegisterServiceAPIConfigs 注册服务API，覆盖更新
	RegisterServiceAPIConfigs(apis []*configpb.ClusterClientApi)
	// GetServiceAPIConfig serviceName统一使用常量定义
	GetServiceAPIConfig(serviceName ServiceName) (*Config, error)
}

type apiManager struct {
	configMap   map[ServiceName]*Config
	configMutex sync.RWMutex
}

// RegisterServiceAPIConfigs 注册服务API，覆盖已有服务
func (s *apiManager) RegisterServiceAPIConfigs(apis []*configpb.ClusterClientApi) {
	s.configMutex.Lock()
	defer s.configMutex.Unlock()
	for i := range apis {
		conf := &Config{}
		conf.SetByPbClusterClientApi(apis[i])
		s.configMap[ServiceName(apis[i].ServiceName)] = conf
	}
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
