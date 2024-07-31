package serverutil

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	configpb "github.com/go-micro-saas/service-kit/api/config"
	apputil "github.com/go-micro-saas/service-kit/app"
	middlewareutil "github.com/go-micro-saas/service-kit/middleware"
	consulapi "github.com/hashicorp/consul/api"
	authpkg "github.com/ikaiguang/go-srv-kit/kratos/auth"
	registrypkg "github.com/ikaiguang/go-srv-kit/kratos/registry"
	pkgerrors "github.com/pkg/errors"
	stdlog "log"
)

// NewApp .
func NewApp(
	appConfig *configpb.App,
	httpConfig *configpb.Server_HTTP,
	grpcConfig *configpb.Server_GRPC,
	logger log.Logger,
	loggerForMiddleware log.Logger,
	authManager authpkg.AuthRepo,
	authWhiteList map[string]middlewareutil.TransportServiceKind,
	settingConfig *configpb.Setting,
	consulClient *consulapi.Client,
) (app *kratos.App, err error) {
	//tracerutil.InitTracer()
	//tracerutil.InitTracerWithJaegerExporter()

	// 服务
	var servers []transport.Server

	// http
	if httpConfig.GetEnable() {
		hs, err := NewHTTPServer(httpConfig, loggerForMiddleware, authManager, authWhiteList)
		if err != nil {
			return app, err
		}
		servers = append(servers, hs)
	}
	if grpcConfig.GetEnable() {
		gs, err := NewGRPCServer(grpcConfig, loggerForMiddleware, authManager, authWhiteList)
		if err != nil {
			return app, err
		}
		servers = append(servers, gs)
	}
	if len(servers) == 0 {
		err = pkgerrors.New("服务列表为空")
		return app, err
	}

	// appid
	appID := apputil.ID(appConfig)
	appConfig.Id = appID
	if appConfig.GetMetadata() == nil {
		appConfig.Metadata = make(map[string]string)
	}
	appConfig.GetMetadata()["id"] = appID

	// app
	var (
		appOptions = []kratos.Option{
			kratos.ID(appID),
			kratos.Name(appID),
			kratos.Version(appConfig.GetServerVersion()),
			kratos.Metadata(appConfig.GetMetadata()),
			kratos.Logger(logger),
			kratos.Server(servers...),
		}
	)

	// 启用服务注册中心
	if settingConfig.GetEnableServiceRegistry() {
		stdlog.Println("|*** 加载：服务注册与发现")

		r, err := registrypkg.NewConsulRegistry(consulClient)
		if err != nil {
			return app, err
		}
		//registrypkg.SetRegistryType(registrypkg.RegistryTypeConsul)
		appOptions = append(appOptions, kratos.Registrar(r))
	}

	// 路由；放置在"服务注册"后，否则 engineHandler.RegistryType 不生效
	//err = routes.RegisterRoutes(hs, gs)
	//if err != nil {
	//	return app, err
	//}

	app = kratos.New(appOptions...)
	return app, err
}
