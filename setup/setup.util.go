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
	return nil
}
