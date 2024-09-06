package clientutil

import (
	"github.com/go-kratos/kratos/v2/transport/http"
	configpb "github.com/go-micro-saas/service-kit/api/config"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	stdgrpc "google.golang.org/grpc"
	"strings"
)

var (
	uninitializedConsulClientError = errorpkg.ErrorBadRequest("uninitialized: consulClient == nil")
	uninitializedEtcdClientError   = errorpkg.ErrorBadRequest("uninitialized: etcdClient == nil")
	uninitializedGRPCConnError     = errorpkg.ErrorBadRequest("uninitialized: grpcConn == nil")
	uninitializedHTTPClientError   = errorpkg.ErrorBadRequest("uninitialized: httpClient == nil")
)

type ServiceAPIManager interface {
	// RegisterServiceAPIConfigs 注册服务API，覆盖更新
	RegisterServiceAPIConfigs(apis []*configpb.ClusterServiceApi, opts ...Option) error
	// GetServiceAPIConfig 获取服务配置
	GetServiceAPIConfig(serviceName ServiceName) (*Config, error)
	// NewAPIConnection 实例化客户端链接
	NewAPIConnection(serviceName ServiceName) (ServiceAPIConnection, error)
}

type ServiceAPIConnection interface {
	GetTransportType() configpb.ClusterServiceApi_TransportType
	IsTransportType(tt configpb.ClusterServiceApi_TransportType) bool
	GetGRPCConnection() (*stdgrpc.ClientConn, error)
	GetHTTPClient() (*http.Client, error)
}

// ServiceName ...
type ServiceName string

func (s ServiceName) String() string {
	return string(s)
}

// Config ...
type Config struct {
	ServiceName   string                                   // 服务名称
	TransportType configpb.ClusterServiceApi_TransportType // 传输协议：http、grpc、...；默认: HTTP
	RegistryType  configpb.ClusterServiceApi_RegistryType  // 注册类型：注册类型：endpoint、consul、...；配置中心配置：${registry_type}；例： Bootstrap.Consul
	ServiceTarget string                                   // 服务目标：endpoint或registry，例：http://127.0.0.1:8899、discovery:///${registry_endpoint}
}

func (s *Config) SetByPbClusterServiceApi(cfg *configpb.ClusterServiceApi) {
	s.ServiceName = cfg.GetServiceName()
	tt := strings.ToLower(cfg.GetTransportType())
	switch tt {
	default:
		s.TransportType = configpb.ClusterServiceApi_TT_HTTP
	case "http":
		s.TransportType = configpb.ClusterServiceApi_TT_HTTP
	case "grpc":
		s.TransportType = configpb.ClusterServiceApi_TT_GRPC
	}
	rt := strings.ToLower(cfg.GetRegistryType())
	switch rt {
	default:
		s.RegistryType = configpb.ClusterServiceApi_RT_ENDPOINT
	case "endpoint":
		s.RegistryType = configpb.ClusterServiceApi_RT_ENDPOINT
	case "consul":
		s.RegistryType = configpb.ClusterServiceApi_RT_CONSUL
	case "etcd":
		s.RegistryType = configpb.ClusterServiceApi_RT_ETCD
	}
	s.ServiceTarget = cfg.GetServiceTarget()
}

func (s *Config) IsConsulRegistry() bool {
	return s.RegistryType == configpb.ClusterServiceApi_RT_CONSUL
}

func (s *Config) IsEtcdRegistry() bool {
	return s.RegistryType == configpb.ClusterServiceApi_RT_ETCD
}

type clientConnection struct {
	transportType configpb.ClusterServiceApi_TransportType
	grpcConn      *stdgrpc.ClientConn
	httpClient    *http.Client
}

func (c *clientConnection) SetTransportType(tt configpb.ClusterServiceApi_TransportType) {
	_, ok := configpb.ClusterServiceApi_TransportType_name[int32(tt)]
	if ok {
		c.transportType = tt
	}
	if c.transportType == configpb.ClusterServiceApi_TT_UNSPECIFIED {
		c.transportType = configpb.ClusterServiceApi_TT_HTTP
	}
}

func (c *clientConnection) GetTransportType() configpb.ClusterServiceApi_TransportType {
	return c.transportType
}

func (c *clientConnection) IsTransportType(tt configpb.ClusterServiceApi_TransportType) bool {
	return c.transportType == tt
}

func (c *clientConnection) IsHTTPTransport() bool {
	return c.transportType == configpb.ClusterServiceApi_TT_HTTP
}

func (c *clientConnection) IsGRCPTransport() bool {
	return c.transportType == configpb.ClusterServiceApi_TT_GRPC
}

func (c *clientConnection) GetGRPCConnection() (*stdgrpc.ClientConn, error) {
	if c.grpcConn == nil {
		return nil, errorpkg.WithStack(uninitializedGRPCConnError)
	}
	return c.grpcConn, nil
}

func (c *clientConnection) GetHTTPClient() (*http.Client, error) {
	if c.httpClient == nil {
		return nil, errorpkg.WithStack(uninitializedHTTPClientError)
	}
	return c.httpClient, nil
}
