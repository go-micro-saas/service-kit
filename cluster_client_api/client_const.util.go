package clientutil

import (
	configpb "github.com/go-micro-saas/service-kit/api/config"
	"strings"
)

type ClientAPIManager interface {
	// RegisterServiceAPIConfigs 注册服务API，覆盖更新
	RegisterServiceAPIConfigs(apis []*configpb.ClusterClientApi) error
}

// ServiceName ...
type ServiceName string

func (s ServiceName) String() string {
	return string(s)
}

// 示例：仅供参考
const (
	PingService   ServiceName = "ping-service"
	NodeidService ServiceName = "nodeid-service"

	FeishuApi      ServiceName = "feishu-openapi"
	DingtalkApi    ServiceName = "dingtalk-openapi"
	DingtalkApiOld ServiceName = "dingtalk-openapi-old"
)

// Config ...
type Config struct {
	ServiceName   string                                  // 服务名称
	TransportType configpb.ClusterClientApi_TransportType // 传输协议：http、grpc、...；默认: HTTP
	RegistryType  configpb.ClusterClientApi_RegistryType  // 注册类型：注册类型：endpoint、consul、...；配置中心配置：${registry_type}；例： Bootstrap.Consul
	ServiceTarget string                                  // 服务目标：endpoint或registry，例：http://127.0.0.1:8899、discovery:///${registry_endpoint}
}

func (s *Config) SetByPbClusterClientApi(cfg *configpb.ClusterClientApi) {
	s.ServiceName = cfg.GetServiceName()
	tt := strings.ToLower(cfg.GetTransportType())
	switch tt {
	default:
		s.TransportType = configpb.ClusterClientApi_TT_HTTP
	case "http":
		s.TransportType = configpb.ClusterClientApi_TT_HTTP
	case "grpc":
		s.TransportType = configpb.ClusterClientApi_TT_GRPC
	}
	rt := strings.ToLower(cfg.GetRegistryType())
	switch rt {
	default:
		s.RegistryType = configpb.ClusterClientApi_RT_ENDPOINT
	case "endpoint":
		s.RegistryType = configpb.ClusterClientApi_RT_ENDPOINT
	case "consul":
		s.RegistryType = configpb.ClusterClientApi_RT_CONSUL
	case "etcd":
		s.RegistryType = configpb.ClusterClientApi_RT_ETCD
	}
	s.ServiceTarget = cfg.GetServiceTarget()
}
