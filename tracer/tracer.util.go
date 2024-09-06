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
	return middlewarepkg.SetTracer(apputil.ID(apputil.ToAppConfig(appConfig)), opts...)
}

func InitTracer(appConfig *configpb.App) error {
	stdlog.Println("|*** LOADING: Tracer: ...")
	return middlewarepkg.SetTracer(apputil.ID(apputil.ToAppConfig(appConfig)))
}
