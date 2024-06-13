package clientutil

import (
	setuputil "github.com/go-micro-saas/service-kit/setup"
	pingservicev1 "github.com/ikaiguang/go-srv-kit/api/ping/v1/services"
)

// NewPingHTTPClient ...
func NewPingHTTPClient(engineHandler setuputil.Launch, serviceName ServiceName) (pingservicev1.SrvPingHTTPClient, error) {
	conn, err := NewHTTPConnection(engineHandler, serviceName)
	if err != nil {
		return nil, err
	}
	return pingservicev1.NewSrvPingHTTPClient(conn), nil
}

// NewUserHTTPClient ...
//func NewUserHTTPClient(engineHandler setuputil.Launch, serviceName ServiceName) (userservicev1.SrvUserHTTPClient, error) {
//	conn, err := NewHTTPConnection(engineHandler, serviceName)
//	if err != nil {
//		return nil, err
//	}
//	return userservicev1.NewSrvUserHTTPClient(conn), nil
//}

// NewAdminHTTPClient ...
//func NewAdminHTTPClient(engineHandler setuputil.Launch, serviceName ServiceName) (adminservicev1.SrvAdminHTTPClient, error) {
//	conn, err := NewHTTPConnection(engineHandler, serviceName)
//	if err != nil {
//		return nil, err
//	}
//	return adminservicev1.NewSrvAdminHTTPClient(conn), nil
//}
