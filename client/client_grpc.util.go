package clientutil

import (
	setuputil "github.com/go-micro-saas/service-kit/setup"
	pingservicev1 "github.com/ikaiguang/go-srv-kit/api/ping/v1/services"
)

// NewPingGRPCClient ...
func NewPingGRPCClient(engineHandler setuputil.Launch, serviceName ServiceName) (pingservicev1.SrvPingClient, error) {
	conn, err := NewGRPCConnection(engineHandler, serviceName)
	if err != nil {
		return nil, err
	}
	return pingservicev1.NewSrvPingClient(conn), nil
}

// NewUserGRPCClient ...
//func NewUserGRPCClient(engineHandler setuputil.Launch, serviceName ServiceName) (userservicev1.SrvUserClient, error) {
//	conn, err := NewGRPCConnection(engineHandler, serviceName)
//	if err != nil {
//		return nil, err
//	}
//	return userservicev1.NewSrvUserClient(conn), nil
//}

// NewAdminGRPCClient ...
//func NewAdminGRPCClient(engineHandler setuputil.Launch, serviceName ServiceName) (adminservicev1.SrvAdminClient, error) {
//	conn, err := NewGRPCConnection(engineHandler, serviceName)
//	if err != nil {
//		return nil, err
//	}
//	return adminservicev1.NewSrvAdminClient(conn), nil
//}
