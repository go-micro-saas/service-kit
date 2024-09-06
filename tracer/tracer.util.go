package tracerutil

import (
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	stdlog "log"

	configpb "github.com/go-micro-saas/service-kit/api/config"
	apputil "github.com/go-micro-saas/service-kit/app"
	middlewarepkg "github.com/ikaiguang/go-srv-kit/kratos/middleware"
)

func InitTracerWithJaegerExporter(appConfig *configpb.App, exp *otlptrace.Exporter) error {
	stdlog.Println("|*** LOADING: Tracer: ...")
	// Create the Jaeger exporter
	var opts = []middlewarepkg.TracerOption{
		middlewarepkg.WithTracerJaegerExporter(exp),
	}
	ac := &apputil.AppConfig{}
	ac.SetByPbApp(appConfig)
	return middlewarepkg.SetTracer(apputil.ID(ac), opts...)
}

func InitTracer(appConfig *configpb.App) error {
	stdlog.Println("|*** LOADING: Tracer: ...")
	ac := &apputil.AppConfig{}
	ac.SetByPbApp(appConfig)
	return middlewarepkg.SetTracer(apputil.ID(ac))
}
