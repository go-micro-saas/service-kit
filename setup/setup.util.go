package setuputil

import (
	configutil "github.com/go-micro-saas/service-kit/config"
	stdlog "log"
)

func Setup(configFilePath string) error {
	// 开始配置
	stdlog.Println("|==================== LOADING PROGRAM CONFIGURATION : START ====================|")
	defer stdlog.Println("|==================== Loading program configuration : END ====================|")
	conf, err := configutil.Loading(configFilePath)
	if err != nil {
		return err
	}
	configutil.SetConfig(conf)

	// 日志工具
	//loggerManager, err := loggerutil.NewSingletonLoggerManager(conf.GetLog(), conf.GetApp())
	//if err != nil {
	//	return err
	//}
	// mysql gorm 数据库
	// postgres gorm 数据库
	// redis 客户端
	// consul 客户端
	// jaeger
	// 雪花算法
	// 设置调试工具
	// 设置日志工具
	// 服务注册
	return nil
}
