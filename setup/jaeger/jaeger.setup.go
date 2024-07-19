package jaegerutil

import (
	configpb "github.com/go-micro-saas/service-kit/api/config"
	apputil "github.com/go-micro-saas/service-kit/app"
	jaegerpkg "github.com/ikaiguang/go-srv-kit/data/jaeger"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	middlewarepkg "github.com/ikaiguang/go-srv-kit/kratos/middleware"
	"go.opentelemetry.io/otel/exporters/jaeger"
	stdlog "log"
	"sync"
)

type jaegerManager struct {
	conf *configpb.Jaeger

	jaegerOnce     sync.Once
	jaegerExporter *jaeger.Exporter
}

type JaegerManager interface {
	GetExporter() (*jaeger.Exporter, error)
}

func NewJaegerManager(conf *configpb.Jaeger) (JaegerManager, error) {
	if conf == nil {
		e := errorpkg.ErrorBadRequest("[请配置服务再启动] config key : jaeger")
		return nil, errorpkg.WithStack(e)
	}
	return &jaegerManager{
		conf: conf,
	}, nil
}

func (s *jaegerManager) GetExporter() (*jaeger.Exporter, error) {

	var err error
	s.jaegerOnce.Do(func() {
		s.jaegerExporter, err = s.loadingJaegerTraceExporter()
		if err != nil {
			s.jaegerOnce = sync.Once{}
		}
	})
	return s.jaegerExporter, err
}

func (s *jaegerManager) loadingJaegerTraceExporter() (*jaeger.Exporter, error) {
	stdlog.Println("|*** 加载：JaegerExporter：...")
	je, err := jaegerpkg.NewJaegerExporter(ToJaegerConfig(s.conf))
	if err != nil {
		e := errorpkg.ErrorInternalError(err.Error())
		return nil, errorpkg.WithStack(e)
	}
	return je, nil
}

// InitTracerProvider trace provider
func (s *jaegerManager) InitTracerProvider(appConfig *configpb.App) error {
	stdlog.Println("|*** 加载：服务追踪：Tracer")
	// Create the Jaeger exporter
	var opts []middlewarepkg.TracerOption
	if s.conf.GetEnable() {
		exp, err := s.GetExporter()
		if err != nil {
			return err
		}
		opts = append(opts, middlewarepkg.WithTracerJaegerExporter(exp))
	}
	return middlewarepkg.SetTracer(apputil.ID(appConfig), opts...)
}

// ToJaegerConfig ...
func ToJaegerConfig(cfg *configpb.Jaeger) *jaegerpkg.Config {
	return &jaegerpkg.Config{
		Endpoint:          cfg.Endpoint,
		WithHttpBasicAuth: cfg.WithHttpBasicAuth,
		Username:          cfg.Username,
		Password:          cfg.Password,
	}
}
