package serverutil

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	middlewareutil "github.com/go-micro-saas/service-kit/middleware"
	setuputil "github.com/go-micro-saas/service-kit/setup"
)

// Services 各个服务注册结果
type Services struct {
}

type Servers struct {
	HTTP *http.Server
	GRPC *grpc.Server
}

func NewServers(launcherManager setuputil.LauncherManager, whiteList map[string]middlewareutil.TransportServiceKind) (*Servers, error) {
	hs, err := NewHTTPServer(launcherManager, whiteList)
	if err != nil {
		return nil, err
	}
	gs, err := NewGRPCServer(launcherManager, whiteList)
	if err != nil {
		return nil, err
	}
	return &Servers{
		HTTP: hs,
		GRPC: gs,
	}, nil
}

func GetHTTPServer(services *Servers) (*http.Server, error) {
	return services.HTTP, nil
}

func GetGRPCServer(services *Servers) (*grpc.Server, error) {
	return services.GRPC, nil
}
