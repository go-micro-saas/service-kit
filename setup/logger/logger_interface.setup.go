package loggerutil

import "io"

type LoggerManager interface {
	GetWriter() (io.Writer, error)
	GetLoggers() (*Loggers, error)
}
