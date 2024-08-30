package serverutil

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	configpb "github.com/go-micro-saas/service-kit/api/config"
	apputil "github.com/go-micro-saas/service-kit/app"
	consulutil "github.com/go-micro-saas/service-kit/consul"
	jaegerutil "github.com/go-micro-saas/service-kit/jaeger"
	setuputil "github.com/go-micro-saas/service-kit/setup"
	tracerutil "github.com/go-micro-saas/service-kit/tracer"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	registrypkg "github.com/ikaiguang/go-srv-kit/kratos/registry"
	stdlog "log"
)

// NewApp .
func NewApp(launcherManager setuputil.LauncherManager, hs *http.Server, gs *grpc.Server) (*kratos.App, error) {
	conf := launcherManager.GetConfig()
	if err := InitTracer(conf); err != nil {
		return nil, err
	}

	// logger
	logger, err := launcherManager.GetLogger()
	if err != nil {
		return nil, err
	}

	// 服务
	var servers []transport.Server

	// http
	httpConfig := conf.GetServer().GetHttp()
	if httpConfig.GetEnable() {
		servers = append(servers, hs)
	}

	// grpc
	grpcConfig := conf.GetServer().GetGrpc()
	if grpcConfig.GetEnable() {
		servers = append(servers, gs)
	}
	if len(servers) == 0 {
		e := errorpkg.ErrorInvalidParameter("server list cannot be empty")
		return nil, errorpkg.WithStack(e)
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
		stdlog.Println("|*** LOADING: ServiceRegistry: ...")
		consulManager, err := consulutil.NewSingletonConsulManager(conf.GetConsul())
		if err != nil {
			return nil, err
		}
		consulClient, err := consulManager.GetClient()
		if err != nil {
			return nil, err
		}
		r, err := registrypkg.NewConsulRegistry(consulClient)
		if err != nil {
			return nil, err
		}
		appOptions = append(appOptions, kratos.Registrar(r))
	}

	return kratos.New(appOptions...), err
}

func InitTracer(conf *configpb.Bootstrap) error {
	if conf.GetSetting().GetEnableJaegerTracer() {
		jaegerManager, err := jaegerutil.NewSingletonJaegerManager(conf.GetJaeger())
		if err != nil {
			return err
		}
		exp, err := jaegerManager.GetExporter()
		if err != nil {
			return err
		}
		return tracerutil.InitTracerWithJaegerExporter(conf.GetApp(), exp)
	}
	return tracerutil.InitTracer(conf.GetApp())
}
