package clientutil

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	configpb "github.com/go-micro-saas/service-kit/api/config"
	apputil "github.com/go-micro-saas/service-kit/app"
	clientpkg "github.com/ikaiguang/go-srv-kit/kratos/client"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	middlewarepkg "github.com/ikaiguang/go-srv-kit/kratos/middleware"
)

func (s *serviceAPIManager) NewHTTPClient(apiConfig *Config, otherOpts ...http.ClientOption) (*http.Client, error) {
	var opts = []http.ClientOption{
		http.WithTimeout(DefaultTimeout),
	}
	opts = append(opts, apputil.ClientDecoderEncoder()...)

	// 中间件
	logHelper := log.NewHelper(s.opt.logger)
	opts = append(opts, http.WithMiddleware(middlewarepkg.DefaultClientMiddlewares(logHelper)...))

	// 服务端点
	endpointOpts, err := s.getHTTPEndpointOptions(apiConfig)
	if err != nil {
		return nil, err
	}
	opts = append(opts, endpointOpts...)
	logHelper.Infow(
		"client.serviceName", apiConfig.ServiceName,
		"client.transportType", apiConfig.TransportType.String(),
		"client.registryType", apiConfig.RegistryType.String(),
		"client.serviceTarget", apiConfig.ServiceTarget,
	)

	// 其他
	opts = append(opts, otherOpts...)

	// http 链接
	conn, err := clientpkg.NewHTTPClient(context.Background(), opts...)
	if err != nil {
		e := errorpkg.ErrorInternalServer(err.Error())
		return nil, errorpkg.WithStack(e)
	}
	return conn, nil
}

// getHTTPEndpointOptions 获取服务端点
func (s *serviceAPIManager) getHTTPEndpointOptions(apiConfig *Config) ([]http.ClientOption, error) {
	var opts []http.ClientOption

	// endpoint
	opts = append(opts, http.WithEndpoint(apiConfig.ServiceTarget))

	// registry
	switch apiConfig.RegistryType {
	case configpb.RegistryTypeEnum_CONSUL, configpb.RegistryTypeEnum_ETCD:
		r, err := s.getRegistryDiscovery(apiConfig)
		if err != nil {
			return nil, err
		}
		opts = append(opts, http.WithDiscovery(r))
	}
	return opts, nil
}
