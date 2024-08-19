package jaegerutil

import (
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"sync"

	configpb "github.com/go-micro-saas/service-kit/api/config"
)

var (
	singletonMutex         sync.Once
	singletonJaegerManager JaegerManager
)

func NewSingletonJaegerManager(conf *configpb.Jaeger) (JaegerManager, error) {
	var err error
	singletonMutex.Do(func() {
		singletonJaegerManager, err = NewJaegerManager(conf)
	})
	if err != nil {
		singletonMutex = sync.Once{}
	}
	return singletonJaegerManager, err
}

func GetJaegerExporter(jaegerManager JaegerManager) (*otlptrace.Exporter, error) {
	return jaegerManager.GetExporter()
}
