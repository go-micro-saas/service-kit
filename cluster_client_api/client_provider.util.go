package clientutil

import (
	pingservicev1 "github.com/ikaiguang/go-srv-kit/api/ping/v1/services"
	"sync"
)

var (
	_apiConnection = sync.Map{}
)

func NewSingletonClientAPIConnection(clientAPIManager ClientAPIManager, serviceName ServiceName) (APIConnection, error) {
	cc, ok := _apiConnection.Load(serviceName)
	if ok {
		if conn, ok := cc.(APIConnection); ok {
			return conn, nil
		}
	}
	conn, err := NewClientAPIConnection(clientAPIManager, serviceName)
	if err != nil {
		return nil, err
	}
	_apiConnection.Store(serviceName, conn)
	return conn, nil
}

func NewClientAPIConnection(clientAPIManager ClientAPIManager, serviceName ServiceName) (APIConnection, error) {
	conn, err := clientAPIManager.NewAPIConnection(serviceName)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// 示例：仅供参考
const (
	PingService   ServiceName = "ping-service"
	NodeidService ServiceName = "nodeid-service"

	FeishuApi      ServiceName = "feishu-openapi"
	DingtalkApi    ServiceName = "dingtalk-openapi"
	DingtalkApiOld ServiceName = "dingtalk-openapi-old"
)

// NewPingGRPCClient ...
func NewPingGRPCClient(clientAPIManager ClientAPIManager, rewriteServiceName ...ServiceName) (pingservicev1.SrvPingClient, error) {
	serviceName := PingService
	for i := range rewriteServiceName {
		serviceName = rewriteServiceName[i]
	}
	conn, err := NewSingletonClientAPIConnection(clientAPIManager, serviceName)
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
func NewPingHTTPClient(clientAPIManager ClientAPIManager, rewriteServiceName ...ServiceName) (pingservicev1.SrvPingHTTPClient, error) {
	serviceName := PingService
	for i := range rewriteServiceName {
		serviceName = rewriteServiceName[i]
	}
	conn, err := NewSingletonClientAPIConnection(clientAPIManager, serviceName)
	if err != nil {
		return nil, err
	}
	httpClient, err := conn.GetHTTPClient()
	if err != nil {
		return nil, err
	}
	return pingservicev1.NewSrvPingHTTPClient(httpClient), nil
}
