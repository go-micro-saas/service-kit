package jaegerutil

import (
	"context"
	stdlog "log"
	"sync"

	configpb "github.com/go-micro-saas/service-kit/api/config"
	jaegerpkg "github.com/ikaiguang/go-srv-kit/data/jaeger"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	"go.opentelemetry.io/otel/exporters/jaeger"
)

type jaegerManager struct {
	conf *configpb.Jaeger

	jaegerOnce     sync.Once
	jaegerExporter *jaeger.Exporter
}

type JaegerManager interface {
	Enable() bool
	GetExporter() (*jaeger.Exporter, error)
	Close() error
}

func NewJaegerManager(conf *configpb.Jaeger) (JaegerManager, error) {
	if conf == nil {
		e := errorpkg.ErrorBadRequest("[CONFIGURATION] config error, key = jaeger")
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

func (s *jaegerManager) Close() error {
	if s.jaegerExporter != nil {
		stdlog.Println("|*** STOP: close: jaegerExporter")
		err := s.jaegerExporter.Shutdown(context.Background())
		if err != nil {
			stdlog.Println("|*** STOP: close: jaegerExporter failed: ", err.Error())
			return err
		}
	}
	return nil
}

func (s *jaegerManager) Enable() bool {
	return s.conf.GetEnable()
}

func (s *jaegerManager) loadingJaegerTraceExporter() (*jaeger.Exporter, error) {
	stdlog.Println("|*** LOADING: JaegerExporter: ...")
	je, err := jaegerpkg.NewJaegerExporter(ToJaegerConfig(s.conf))
	if err != nil {
		e := errorpkg.ErrorInternalError(err.Error())
		return nil, errorpkg.WithStack(e)
	}
	return je, nil
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
