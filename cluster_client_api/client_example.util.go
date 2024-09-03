package clientutil

//import (
//	setuputil "github.com/go-micro-saas/service-kit/setup"
//	pingservicev1 "github.com/ikaiguang/go-srv-kit/api/ping/v1/services"
//)
//
//// NewPingGRPCClient ...
//func NewPingGRPCClient(launcherManager setuputil.LauncherManager, serviceName ServiceName) (pingservicev1.SrvPingClient, error) {
//	conn, err := NewGRPCConnection(launcherManager, serviceName)
//	if err != nil {
//		return nil, err
//	}
//	return pingservicev1.NewSrvPingClient(conn), nil
//}
//
//// NewPingHTTPClient ...
//func NewPingHTTPClient(launcherManager setuputil.LauncherManager, serviceName ServiceName) (pingservicev1.SrvPingHTTPClient, error) {
//	conn, err := NewHTTPConnection(launcherManager, serviceName)
//	if err != nil {
//		return nil, err
//	}
//	return pingservicev1.NewSrvPingHTTPClient(conn), nil
//}
