package serverutil

import (
	middlewareutil "github.com/go-micro-saas/service-kit/middleware"
	pingservicev1 "github.com/ikaiguang/go-srv-kit/api/ping/v1/services"
)

// GetAuthWhiteList 验证白名单
func GetAuthWhiteList() map[string]middlewareutil.TransportServiceKind {
	// 白名单
	whiteList := make(map[string]middlewareutil.TransportServiceKind)

	// 测试
	whiteList[pingservicev1.OperationSrvPingPing] = middlewareutil.TransportServiceKindALL
	whiteList["/ws/v1/testdata/websocket"] = middlewareutil.TransportServiceKindALL

	return whiteList
}
