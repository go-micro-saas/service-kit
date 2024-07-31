package serverutil

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	configpb "github.com/go-micro-saas/service-kit/api/config"
	apputil "github.com/go-micro-saas/service-kit/app"
	consulutil "github.com/go-micro-saas/service-kit/consul"
	jaegerutil "github.com/go-micro-saas/service-kit/jaeger"
	middlewareutil "github.com/go-micro-saas/service-kit/middleware"
	tracerutil "github.com/go-micro-saas/service-kit/tracer"
	authpkg "github.com/ikaiguang/go-srv-kit/kratos/auth"
	registrypkg "github.com/ikaiguang/go-srv-kit/kratos/registry"
	pkgerrors "github.com/pkg/errors"
	stdlog "log"
)

func InitTracer(conf *configpb.Bootstrap) error {
	if conf.GetSetting().GetEnableJaegerTracer() {
		jaegerManager, err := jaegerutil.NewSingletonJaegerManager(conf.GetJaeger())
		if err != nil {
			return err
		}
		return tracerutil.InitTracerWithJaegerExporter(conf.GetApp(), jaegerManager)
	}
	return tracerutil.InitTracer(conf.GetApp())
}

// NewApp .
func NewApp(
	conf *configpb.Bootstrap,
	logger log.Logger,
	loggerForMiddleware log.Logger,
	authManager authpkg.AuthRepo,
	authWhiteList map[string]middlewareutil.TransportServiceKind,
) (app *kratos.App, err error) {
	if err := InitTracer(conf); err != nil {
		return app, err
	}

	// 服务
	var servers []transport.Server

	// http
	httpConfig := conf.GetServer().GetHttp()
	if httpConfig.GetEnable() {
		hs, err := NewHTTPServer(httpConfig, loggerForMiddleware, authManager, authWhiteList)
		if err != nil {
			return app, err
		}
		servers = append(servers, hs)
	}

	// grpc
	grpcConfig := conf.GetServer().GetGrpc()
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
	appConfig := conf.GetApp()
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
	if conf.GetSetting().GetEnableConsulRegistry() {
		stdlog.Println("|*** 加载：服务注册与发现")
		consulManager, err := consulutil.NewSingletonConsulManager(conf.GetConsul())
		if err != nil {
			return app, err
		}
		consulClient, err := consulManager.GetClient()
		if err != nil {
			return app, err
		}
		r, err := registrypkg.NewConsulRegistry(consulClient)
		if err != nil {
			return app, err
		}
		appOptions = append(appOptions, kratos.Registrar(r))
	}

	app = kratos.New(appOptions...)
	return app, err
}
