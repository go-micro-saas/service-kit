package clientutil

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	configpb "github.com/go-micro-saas/service-kit/api/config"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	middlewarepkg "github.com/ikaiguang/go-srv-kit/kratos/middleware"
	stdgrpc "google.golang.org/grpc"
	"time"
)

const (
	DefaultTimeout = time.Minute
)

func (s *serviceAPIManager) NewGRPCConnection(apiConfig *Config, otherOpts ...grpc.ClientOption) (*stdgrpc.ClientConn, error) {
	var opts = []grpc.ClientOption{
		grpc.WithTimeout(DefaultTimeout),
		grpc.WithHealthCheck(true),
		grpc.WithPrintDiscoveryDebugLog(true),
		//grpc.WithOptions(stdgrpc.WithTimeout(DefaultTimeout)), // dial option
	}

	// 中间件
	logHelper := log.NewHelper(s.opt.logger)
	opts = append(opts, grpc.WithMiddleware(middlewarepkg.DefaultClientMiddlewares(logHelper)...))

	// 服务端点
	endpointOpts, err := s.getGRPCEndpointOptions(apiConfig)
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

	// grpc 链接
	conn, err := grpc.DialInsecure(context.Background(), opts...)
	if err != nil {
		e := errorpkg.ErrorInternalServer(err.Error())
		return nil, errorpkg.WithStack(e)
	}
	return conn, nil
}

func (s *serviceAPIManager) getGRPCEndpointOptions(apiConfig *Config) ([]grpc.ClientOption, error) {
	var opts []grpc.ClientOption

	// endpoint
	opts = append(opts, grpc.WithEndpoint(apiConfig.ServiceTarget))

	// registry
	switch apiConfig.RegistryType {
	case configpb.RegistryTypeEnum_CONSUL, configpb.RegistryTypeEnum_ETCD:
		r, err := s.getRegistryDiscovery(apiConfig)
		if err != nil {
			return nil, err
		}
		opts = append(opts, grpc.WithDiscovery(r))
	}
	return opts, nil
}