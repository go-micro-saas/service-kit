package clientutil

import (
	"github.com/go-kratos/kratos/v2/transport/http"
	pingservicev1 "github.com/ikaiguang/go-srv-kit/api/ping/v1/services"
	stdgrpc "google.golang.org/grpc"
)

// 示例：仅供参考
const (
	PingService   ServiceName = "ping-service"
	NodeidService ServiceName = "nodeid-service"

	FeishuApi      ServiceName = "feishu-openapi"
	DingtalkApi    ServiceName = "dingtalk-openapi"
	DingtalkApiOld ServiceName = "dingtalk-openapi-old"
)

// NewPingGRPCClient ...
func NewPingGRPCClient(apiManager ClientAPIManager) (pingservicev1.SrvPingClient, error) {
	conf, err := apiManager.GetServiceAPIConfig(PingService)
	if err != nil {
		return nil, err
	}
	_ = conf
	var conn *stdgrpc.ClientConn
	//if err != nil {
	//	return nil, err
	//}
	return pingservicev1.NewSrvPingClient(conn), nil
}

// NewPingHTTPClient ...
func NewPingHTTPClient(apiManager ClientAPIManager) (pingservicev1.SrvPingHTTPClient, error) {
	conf, err := apiManager.GetServiceAPIConfig(PingService)
	if err != nil {
		return nil, err
	}
	_ = conf
	var conn *http.Client
	//if err != nil {
	//	return nil, err
	//}
	return pingservicev1.NewSrvPingHTTPClient(conn), nil
}
