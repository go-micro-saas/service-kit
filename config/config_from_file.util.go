package configutil

import (
	stdlog "log"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	configpb "github.com/go-micro-saas/service-kit/api/config"
	apputil "github.com/go-micro-saas/service-kit/app"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
)

func LoadingFile(filePath string) (*configpb.Bootstrap, error) {
	stdlog.Println("|==================== LOADING FILE CONFIGURATION : START ====================|")
	defer stdlog.Println()
	defer stdlog.Println("|==================== LOADING FILE CONFIGURATION : END ====================|")

	p, err := apputil.RuntimePath()
	if err != nil {
		return nil, err
	}
	stdlog.Println("|*** INFO: program running path: ", p)

	var opts []config.Option
	stdlog.Println("|*** LOADING: file configuration path: ", filePath)
	opts = append(opts, config.WithSource(file.NewSource(filePath)))

	handler := config.New(opts...)
	defer func() {
		stdlog.Println("|*** LOADING: COMPLETE : file configuration path: ", filePath)
		_ = handler.Close()
	}()

	// 加载配置
	if err = handler.Load(); err != nil {
		err = errorpkg.WithStack(errorpkg.ErrorInternalError(err.Error()))
		return nil, err
	}

	// 读取配置文件
	conf := &configpb.Bootstrap{}
	if err = handler.Scan(conf); err != nil {
		err = errorpkg.WithStack(errorpkg.ErrorInternalError(err.Error()))
		return nil, err
	}
	return conf, nil
}
