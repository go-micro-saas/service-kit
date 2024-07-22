package tracerutil

import (
	stdlog "log"

	configpb "github.com/go-micro-saas/service-kit/api/config"
	apputil "github.com/go-micro-saas/service-kit/app"
	jaegerutil "github.com/go-micro-saas/service-kit/jaeger"
	middlewarepkg "github.com/ikaiguang/go-srv-kit/kratos/middleware"
)

func InitTracerWithJaegerExporter(appConfig *configpb.App, manager jaegerutil.JaegerManager) error {
	stdlog.Println("|*** LOADING: Tracer: ...")
	// Create the Jaeger exporter
	var opts []middlewarepkg.TracerOption
	if manager.Enable() {
		exp, err := manager.GetExporter()
		if err != nil {
			return err
		}
		opts = append(opts, middlewarepkg.WithTracerJaegerExporter(exp))
	}
	return middlewarepkg.SetTracer(apputil.ID(appConfig), opts...)
}

func InitTracer(appConfig *configpb.App) error {
	stdlog.Println("|*** LOADING: Tracer: ...")
	return middlewarepkg.SetTracer(apputil.ID(appConfig))
}
