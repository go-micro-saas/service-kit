package setuputil

import "io"

type LoggerManager interface {
	GetWriter() (io.Writer, error)
}
