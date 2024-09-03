package clientutil

import (
	configpb "github.com/go-micro-saas/service-kit/api/config"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	"strings"
	"sync"
)

var (
	_serviceConfigMap   = map[ServiceName]*Config{}
	_serviceConfigMutex = sync.RWMutex{}
)

func RegisterClientAPIConfig(apis []*configpb.ClusterClientApi) {
	_serviceConfigMutex.Lock()
	defer _serviceConfigMutex.Unlock()
	for i := range apis {
		conf := &Config{}
		conf.SetByPbClusterClientApi(apis[i])
		_serviceConfigMap[ServiceName(apis[i].ServiceName)] = conf
	}
}

func GetClientAPIConfig(serviceName ServiceName) (*Config, error) {
	if serviceName.String() == "" {
		e := errorpkg.ErrorBadRequest("service name cannot be empty")
		return nil, errorpkg.WithStack(e)
	}
	_serviceConfigMutex.RLock()
	defer _serviceConfigMutex.RUnlock()
	conf, ok := _serviceConfigMap[serviceName]
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

// 示例：仅供参考
const (
	PingService   ServiceName = "ping-service"
	NodeidService ServiceName = "nodeid-service"

	FeishuApi      ServiceName = "feishu-openapi"
	DingtalkApi    ServiceName = "dingtalk-openapi"
	DingtalkApiOld ServiceName = "dingtalk-openapi-old"
)

// ServiceName ...
type ServiceName string

func (s ServiceName) String() string {
	return string(s)
}

// Config ...
type Config struct {
	ServiceName   string                                  // 服务名称
	TransportType configpb.ClusterClientApi_TransportType // 传输协议：http、grpc、...；默认: HTTP
	RegistryType  configpb.ClusterClientApi_RegistryType  // 注册类型：注册类型：endpoint、consul、etcd、...；配置中心配置：${registry_type}；例： Bootstrap.Consul
	ServiceTarget string                                  // 服务目标：endpoint或registry，例：http://127.0.0.1:8899、discovery:///${registry_endpoint}
}

func (s *Config) SetByPbClusterClientApi(cfg *configpb.ClusterClientApi) {
	s.ServiceName = cfg.ServiceName
	tt := strings.ToLower(cfg.TransportType)
	switch tt {
	default:
		s.TransportType = configpb.ClusterClientApi_HTTP
	case "http":
		s.TransportType = configpb.ClusterClientApi_HTTP
	case "grpc":
		s.TransportType = configpb.ClusterClientApi_GRPC
	}
	rt := strings.ToLower(cfg.RegistryType)
	switch rt {
	default:
		s.RegistryType = configpb.ClusterClientApi_ENDPOINT
	case configpb.ClusterClientApi_ENDPOINT:
		s.RegistryType = configpb.ClusterClientApi_ENDPOINT
	case configpb.ClusterClientApi_CONSUL:
		s.RegistryType = configpb.ClusterClientApi_CONSUL
	case configpb.ClusterClientApi_ETCD:
		s.RegistryType = configpb.ClusterClientApi_ETCD
	}
	s.ServiceTarget = cfg.ServiceTarget
}
