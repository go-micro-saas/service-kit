package exportservices

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	configutil "github.com/go-micro-saas/service-kit/config"
	serverutil "github.com/go-micro-saas/service-kit/server"
	setuputil "github.com/go-micro-saas/service-kit/setup"
	"github.com/go-micro-saas/service-kit/testdata/ping-service/internal/conf"
)

func ExportServiceConfig() []configutil.Option {
	return conf.LoadServiceConfig()
}

func ExportServices(launcherManager setuputil.LauncherManager, hs *http.Server, gs *grpc.Server) (*serverutil.Services, func(), error) {
	return exportServices(launcherManager, hs, gs)
}
