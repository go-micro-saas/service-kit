package clientutil

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	curlpkg "github.com/ikaiguang/go-srv-kit/kit/curl"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	middlewarepkg "github.com/ikaiguang/go-srv-kit/kratos/middleware"
	stdgrpc "google.golang.org/grpc"
)

const (
	defaultTimeout = curlpkg.DefaultTimeout
)

func (s *apiManager) NewGRPCConnection(logger log.Logger, serviceName ServiceName, otherOpts ...grpc.ClientOption) (*stdgrpc.ClientConn, error) {
	var opts = []grpc.ClientOption{
		grpc.WithTimeout(defaultTimeout),
		grpc.WithHealthCheck(true),
		grpc.WithPrintDiscoveryDebugLog(true),
		//grpc.WithOptions(stdgrpc.WithTimeout(defaultTimeout)), // dial option
	}

	// 中间件
	logHelper := log.NewHelper(logger)
	opts = append(opts, grpc.WithMiddleware(middlewarepkg.DefaultClientMiddlewares(logHelper)...))

	// 服务端点
	endpointOpts, err := s.getGRPCEndpoint(serviceName)
	if err != nil {
		return nil, err
	}
	opts = append(opts, endpointOpts...)

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

// getGRPCEndpoint 获取服务端点
func (s *apiManager) getGRPCEndpoint(serviceName ServiceName) ([]grpc.ClientOption, error) {
	//apiConfig, err := s.GetServiceAPIConfig(serviceName)
	//if err != nil {
	//	return nil, err
	//}
	//
	//var (
	//	clientKind = transport.KindGRPC
	//	opts       []grpc.ClientOption
	//	discovery  registry.Discovery
	//	endpoint   string
	//)
	//switch registryType {
	//case registrypkg.RegistryTypeConsul:
	//	discovery, endpoint, err = getRegistryAndServerEndpoint(engineHandler, serviceName, endpointInfo.RegistryName)
	//	if err != nil {
	//		return nil, err
	//	}
	//	opts = append(opts, grpc.WithDiscovery(discovery))
	//default:
	//	endpoint = endpointInfo.GrpcHost
	//}
	//logpkg.Infow(
	//	"client.kind", clientKind,
	//	"client.serviceName", serviceName,
	//	"client.registryType", apiConfig.ServiceTarget,
	//	"client.endpoint", endpoint,
	//)
	//opts = append(opts, grpc.WithEndpoint(endpoint))
	//return opts, nil
	return nil, nil
}
