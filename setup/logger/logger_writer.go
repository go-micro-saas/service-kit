package setuputil

import (
	"io"
	stdlog "log"

	writerpkg "github.com/ikaiguang/go-srv-kit/kit/writer"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
)

func (s *loggerManager) GetWriter() (io.Writer, error) {
	var err error
	s.writerOnce.Do(func() {
		s.writer, err = s.getWriter()
	})
	return s.writer, err
}

func (s *loggerManager) getWriter() (io.Writer, error) {
	fileLoggerConfig := s.conf.GetFile()
	if fileLoggerConfig == nil || !fileLoggerConfig.GetEnable() {
		stdlog.Println("|*** 加载：日志工具：虚拟的文件写手柄")
		writer, err := writerpkg.NewDummyWriter()
		if err != nil {
			e := errorpkg.ErrorInternalError(err.Error())
			return nil, errorpkg.WithStack(e)
		}
		return writer, nil
	}

	// rotate write
	rotateConfig := &writerpkg.ConfigRotate{
		Dir:            fileLoggerConfig.Dir,
		Filename:       fileLoggerConfig.Filename,
		RotateTime:     fileLoggerConfig.RotateTime.AsDuration(),
		RotateSize:     fileLoggerConfig.RotateSize,
		StorageCounter: uint(fileLoggerConfig.StorageCounter),
		StorageAge:     fileLoggerConfig.StorageAge.AsDuration(),
	}
	writer, err := writerpkg.NewRotateFile(rotateConfig)
	if err != nil {
		e := errorpkg.ErrorInternalError(err.Error())
		return nil, errorpkg.WithStack(e)
	}
	return writer, nil
}
