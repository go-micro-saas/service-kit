package setuputil

import (
	configpb "github.com/go-micro-saas/service-kit/api/config"
	configutil "github.com/go-micro-saas/service-kit/config"
)

func LoadingConfig(configFilePath string) (*configpb.Bootstrap, error) {
	conf, err := configutil.Loading(configFilePath)
	if err != nil {
		return nil, err
	}
	configutil.SetConfig(conf)
	return conf, nil
}
