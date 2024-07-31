package serverutil

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/transport/http"
	configpb "github.com/go-micro-saas/service-kit/api/config"
	apputil "github.com/go-micro-saas/service-kit/app"
	middlewareutil "github.com/go-micro-saas/service-kit/middleware"
	authpkg "github.com/ikaiguang/go-srv-kit/kratos/auth"
	stdlog "log"
)

var _ metadata.Option

// NewHTTPServer new HTTP server.
func NewHTTPServer(
	httpConfig *configpb.Server_HTTP,
	loggerForMiddleware log.Logger,
	authManager authpkg.AuthRepo,
	authWhiteList map[string]middlewareutil.TransportServiceKind,
) (srv *http.Server, err error) {
	stdlog.Printf("|*** 加载：HTTP服务：%s\n", httpConfig.GetAddr())

	// options
	var opts []http.ServerOption
	//var opts = []http.ServerOption{
	//	http.Filter(middlewareutil.NewCORS()),
	//}
	if httpConfig.Network != "" {
		opts = append(opts, http.Network(httpConfig.GetNetwork()))
	}
	if httpConfig.Addr != "" {
		opts = append(opts, http.Address(httpConfig.GetAddr()))
	}
	if httpConfig.Timeout != nil {
		opts = append(opts, http.Timeout(httpConfig.GetTimeout().AsDuration()))
	}

	// 编码 与 解码
	opts = append(opts, apputil.ServerDecoderEncoder()...)

	// ===== 中间件 =====
	var (
		logHelper       = log.NewHelper(loggerForMiddleware)
		middlewareSlice = middlewareutil.DefaultServerMiddlewares(logHelper)
	)
	// jwt
	stdlog.Println("|*** 加载：JWT中间件：HTTP")
	jwtMiddleware, err := middlewareutil.NewJWTMiddleware(authManager, authWhiteList)
	if err != nil {
		return srv, err
	}
	middlewareSlice = append(middlewareSlice, jwtMiddleware)

	// 中间件选项
	opts = append(opts, http.Middleware(middlewareSlice...))

	// 服务
	srv = http.NewServer(opts...)
	//v1.RegisterGreeterHTTPServer(srv, greeter)

	return srv, err
}

// RegisterHTTPRoute 注册路由
func RegisterHTTPRoute(srv *http.Server) (err error) {
	stdlog.Println("|*** 注册HTTP路由：...")
	return err
}
