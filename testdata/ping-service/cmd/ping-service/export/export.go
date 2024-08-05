package pingexport

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	setuputil "github.com/go-micro-saas/service-kit/setup"
)

type Services struct {
	HTTP *http.Server
	GRPC *grpc.Server
}

func NewServices(hs *http.Server, gs *grpc.Server) (*Services, error) {
	return &Services{
		HTTP: hs,
		GRPC: gs,
	}, nil
}

func Export(launcherManager setuputil.LauncherManager) (*Services, error) {
	return &Services{
		HTTP: nil,
		GRPC: nil,
	}, nil
}
