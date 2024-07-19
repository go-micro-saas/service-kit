package setuputil

import (
	"io"
	"sync"

	configpb "github.com/go-micro-saas/service-kit/api/config"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
)

type loggerManager struct {
	conf *configpb.Log

	writer     io.Writer
	writerOnce sync.Once
}

func NewLoggerManager(conf *configpb.Log) (LoggerManager, error) {
	if conf == nil {
		e := errorpkg.ErrorBadRequest("[请配置服务再启动] config key : log")
		return nil, errorpkg.WithStack(e)
	}
	return &loggerManager{conf: conf}, nil
}
