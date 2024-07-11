package clientutil

import (
	setuputil "github.com/go-micro-saas/service-kit/mytest/backup/setup"
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
