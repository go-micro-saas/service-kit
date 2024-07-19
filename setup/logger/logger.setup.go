package loggerutil

import (
	stderrors "errors"
	"io"
	stdlog "log"
	"sync"

	"github.com/go-kratos/kratos/v2/log"
	configpb "github.com/go-micro-saas/service-kit/api/config"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
)

type loggerManager struct {
	appConfig *configpb.App
	conf      *configpb.Log

	// 不要直接使用 s.writer, 请使用 GetWriter()
	writer     io.Writer
	writerOnce sync.Once

	// 不要直接使用 s.loggerXxx, 请使用 GetLoggers()
	loggerOnce          sync.Once
	logger              log.Logger
	loggerForMiddleware log.Logger
	loggerForHelper     log.Logger
	loggerCloser        io.Closer
}

type Loggers struct {
	Logger              log.Logger
	LoggerForMiddleware log.Logger
	LoggerForHelper     log.Logger
}

func NewLoggerManager(appConfig *configpb.App, conf *configpb.Log) (LoggerManager, error) {
	if appConfig == nil {
		e := errorpkg.ErrorBadRequest("[请配置服务再启动] config key : app")
		return nil, errorpkg.WithStack(e)
	} else if conf == nil {
		e := errorpkg.ErrorBadRequest("[请配置服务再启动] config key : log")
		return nil, errorpkg.WithStack(e)
	}
	return &loggerManager{
		appConfig: appConfig,
		conf:      conf,
	}, nil
}

func (s *loggerManager) Close() error {
	var errs []error

	// loggers
	if s.loggerCloser != nil {
		stdlog.Println("|*** 退出程序：关闭日志 Logger")
		if err := s.loggerCloser.Close(); err != nil {
			stdlog.Println("|*** 退出程序：关闭日志 Logger 异常：", err.Error())
			errs = append(errs, err)
		}
	}

	// writer
	if s.writer != nil {
		if writerCloser, ok := s.writer.(io.Closer); ok {
			stdlog.Println("|*** 退出程序：关闭日志 Writer")
			if err := writerCloser.Close(); err != nil {
				stdlog.Println("|*** 退出程序：关闭日志 Writer 异常：", err.Error())
				errs = append(errs, err)
			}
		}
	}

	if len(errs) > 0 {
		return stderrors.Join(errs...)
	}
	return nil
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
