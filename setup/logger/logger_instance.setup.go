package setuputil

import (
	"io"
	"sync"

	"github.com/go-kratos/kratos/v2/log"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	logpkg "github.com/ikaiguang/go-srv-kit/kratos/log"
)

func (s *loggerManager) setupLoggerOnce() error {
	var err error
	s.loggerOnce.Do(func() {
		err = s.setupLogger()
		if err != nil {
			s.loggerOnce = sync.Once{}
		}
	})
	return err
}

func (s *loggerManager) setupLogger() error {
	cleanup := &closer{
		cs: make([]io.Closer, 0, 6),
	}

	// logger
	loggerSkip := logpkg.CallerSkipForLogger
	logger, loggerClosers, err := s.loadingLoggerWithCallerSkip(loggerSkip)
	if err != nil {
		return err
	}
	for i := range loggerClosers {
		cleanup.cs = append(cleanup.cs, loggerClosers[i])
	}

	// for middleware
	middlewareSkip := logpkg.CallerSkipForMiddleware
	loggerForMiddleware, loggerClosers, err := s.loadingLoggerWithCallerSkip(middlewareSkip)
	if err != nil {
		return err
	}
	for i := range loggerClosers {
		cleanup.cs = append(cleanup.cs, loggerClosers[i])
	}

	// for helper
	helperSkip := logpkg.CallerSkipForHelper
	loggerForHelper, loggerClosers, err := s.loadingLoggerWithCallerSkip(helperSkip)
	for i := range loggerClosers {
		cleanup.cs = append(cleanup.cs, loggerClosers[i])
	}

	s.logger = logger
	s.loggerForMiddleware = loggerForMiddleware
	s.loggerForHelper = loggerForHelper
	s.loggerCloser = cleanup
	return nil
}

func (s *loggerManager) loadingLoggerWithCallerSkip(skip int) (logger log.Logger, closeFnSlice []io.Closer, err error) {
	// loggers
	var loggers []log.Logger

	// DummyLogger
	stdLogger, err := logpkg.NewDummyLogger()
	if err != nil {
		e := errorpkg.ErrorInternalError(err.Error())
		return logger, closeFnSlice, errorpkg.WithStack(e)
	}

	// 日志 输出到控制台
	consoleLoggerConfig := s.conf.GetConsole()
	if consoleLoggerConfig != nil && consoleLoggerConfig.GetEnable() {
		stdLoggerConfig := &logpkg.ConfigStd{
			Level:      logpkg.ParseLevel(consoleLoggerConfig.GetLevel()),
			CallerSkip: skip,
		}
		stdLoggerImpl, err := logpkg.NewStdLogger(stdLoggerConfig)
		if err != nil {
			e := errorpkg.ErrorInternalError(err.Error())
			return logger, closeFnSlice, errorpkg.WithStack(e)
		}
		closeFnSlice = append(closeFnSlice, stdLoggerImpl)
		stdLogger = stdLoggerImpl
	}
	// 覆盖 stdLogger
	loggers = append(loggers, stdLogger)

	// 日志 输出到文件
	fileLoggerConfig := s.conf.GetFile()
	if fileLoggerConfig != nil && fileLoggerConfig.GetEnable() {
		// file logger
		fileLoggerConfig := &logpkg.ConfigFile{
			Level:      logpkg.ParseLevel(fileLoggerConfig.GetLevel()),
			CallerSkip: skip,

			Dir:      fileLoggerConfig.GetDir(),
			Filename: fileLoggerConfig.GetFilename(),

			RotateTime: fileLoggerConfig.GetRotateTime().AsDuration(),
			RotateSize: fileLoggerConfig.GetRotateSize(),

			StorageCounter: uint(fileLoggerConfig.GetStorageCounter()),
			StorageAge:     fileLoggerConfig.GetStorageAge().AsDuration(),
		}
		writer, err := s.GetWriter()
		if err != nil {
			return logger, closeFnSlice, err
		}
		fileLogger, err := logpkg.NewFileLogger(fileLoggerConfig, logpkg.WithWriter(writer))
		closeFnSlice = append(closeFnSlice, fileLogger)
		if err != nil {
			e := errorpkg.ErrorInternalError(err.Error())
			return logger, closeFnSlice, errorpkg.WithStack(e)
		}
		loggers = append(loggers, fileLogger)
	}

	// 日志工具
	return logpkg.NewMultiLogger(loggers...), closeFnSlice, err
}