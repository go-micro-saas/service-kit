package clientutil

import pingservicev1 "github.com/go-micro-saas/service-kit/testdata/ping-service/api/ping-service/v1/services"

// 示例：仅供参考
const (
	PingService   ServiceName = "ping-service"
	NodeidService ServiceName = "nodeid-service"

	FeishuApi      ServiceName = "feishu-openapi"
	DingtalkApi    ServiceName = "dingtalk-openapi"
	DingtalkApiOld ServiceName = "dingtalk-openapi-old"
)

// NewPingGRPCClient ...
func NewPingGRPCClient(serviceAPIManager ServiceAPIManager, rewriteServiceName ...ServiceName) (pingservicev1.SrvPingClient, error) {
	serviceName := PingService
	for i := range rewriteServiceName {
		serviceName = rewriteServiceName[i]
	}
	conn, err := NewSingletonServiceAPIConnection(serviceAPIManager, serviceName)
	//conn, err := NewServiceAPIConnection(serviceAPIManager, serviceName)
	if err != nil {
		return nil, err
	}
	grpcConn, err := conn.GetGRPCConnection()
	if err != nil {
		return nil, err
	}
	return pingservicev1.NewSrvPingClient(grpcConn), nil
}

// NewPingHTTPClient ...
func NewPingHTTPClient(serviceAPIManager ServiceAPIManager, rewriteServiceName ...ServiceName) (pingservicev1.SrvPingHTTPClient, error) {
	serviceName := PingService
	for i := range rewriteServiceName {
		serviceName = rewriteServiceName[i]
	}
	conn, err := NewSingletonServiceAPIConnection(serviceAPIManager, serviceName)
	//conn, err := NewServiceAPIConnection(serviceAPIManager, serviceName)
	if err != nil {
		return nil, err
	}
	httpClient, err := conn.GetHTTPClient()
	if err != nil {
		return nil, err
	}
	return pingservicev1.NewSrvPingHTTPClient(httpClient), nil
}
