package serverutil

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	configpb "github.com/go-micro-saas/service-kit/api/config"
	middlewareutil "github.com/go-micro-saas/service-kit/middleware"
	authpkg "github.com/ikaiguang/go-srv-kit/kratos/auth"
	stdlog "log"
)

var _ metadata.Option

// NewGRPCServer new a gRPC server.
func NewGRPCServer(
	grpcConfig *configpb.Server_GRPC,
	loggerForMiddleware log.Logger,
	authManager authpkg.AuthRepo,
	authWhiteList map[string]middlewareutil.TransportServiceKind,
) (srv *grpc.Server, err error) {
	stdlog.Printf("|*** 加载：GRPC服务：%s\n", grpcConfig.GetAddr())

	// options
	var opts []grpc.ServerOption
	if grpcConfig.Network != "" {
		opts = append(opts, grpc.Network(grpcConfig.GetNetwork()))
	}
	if grpcConfig.Addr != "" {
		opts = append(opts, grpc.Address(grpcConfig.GetAddr()))
	}
	if grpcConfig.Timeout != nil {
		opts = append(opts, grpc.Timeout(grpcConfig.GetTimeout().AsDuration()))
	}

	// ===== 中间件 =====
	var (
		logHelper       = log.NewHelper(loggerForMiddleware)
		middlewareSlice = middlewareutil.DefaultServerMiddlewares(logHelper)
	)
	// jwt
	stdlog.Println("|*** 加载：JWT中间件：GRPC")
	jwtMiddleware, err := middlewareutil.NewJWTMiddleware(authManager, authWhiteList)
	if err != nil {
		return srv, err
	}
	middlewareSlice = append(middlewareSlice, jwtMiddleware)

	// 中间件选项
	opts = append(opts, grpc.Middleware(middlewareSlice...))

	// 服务
	srv = grpc.NewServer(opts...)
	//v1.RegisterGreeterServer(srv, greeter)

	return srv, err
}

// RegisterGRPCRoute 注册路由
func RegisterGRPCRoute(srv *grpc.Server) (err error) {
	stdlog.Println("|*** 注册GRPC路由：...")
	return err
}
