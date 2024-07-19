package setuputil

import (
	stderrors "errors"
	"io"
	"sync"

	"github.com/go-kratos/kratos/v2/log"
	configpb "github.com/go-micro-saas/service-kit/api/config"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
)

type loggerManager struct {
	conf *configpb.Log

	writer     io.Writer
	writerOnce sync.Once

	loggerOnce          sync.Once
	logger              log.Logger
	loggerForMiddleware log.Logger
	loggerForHelper     log.Logger
	loggerCloser        io.Closer
}

func NewLoggerManager(conf *configpb.Log) (LoggerManager, error) {
	if conf == nil {
		e := errorpkg.ErrorBadRequest("[请配置服务再启动] config key : log")
		return nil, errorpkg.WithStack(e)
	}
	return &loggerManager{conf: conf}, nil
}

type closer struct {
	cs []io.Closer
}

func (c *closer) Close() error {
	var errs []error
	for _, v := range c.cs {
		if err := v.Close(); err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return stderrors.Join(errs...)
	}
	return nil
}
