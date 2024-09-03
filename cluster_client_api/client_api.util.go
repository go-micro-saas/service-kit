package clientutil

import (
	configpb "github.com/go-micro-saas/service-kit/api/config"
	"strings"
)

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
