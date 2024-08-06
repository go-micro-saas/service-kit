package runservices

import (
	"github.com/go-kratos/kratos/v2"
	configutil "github.com/go-micro-saas/service-kit/config"
)

func GetServerApp(configFilePath string, configOpts ...configutil.Option) (*kratos.App, func(), error) {
	return initServiceApp(configFilePath, configOpts...)
}
