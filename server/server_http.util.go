package serverutil

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/transport/http"
	apputil "github.com/go-micro-saas/service-kit/app"
	configutil "github.com/go-micro-saas/service-kit/config"
	middlewareutil "github.com/go-micro-saas/service-kit/middleware"
	setuputil "github.com/go-micro-saas/service-kit/setup"
	stdlog "log"
)

var _ metadata.Option

// NewHTTPServer new HTTP server.
func NewHTTPServer(
	launcherManager setuputil.LauncherManager,
	authWhiteList map[string]middlewareutil.TransportServiceKind,
) (srv *http.Server, err error) {
	httpConfig := configutil.HTTPConfig(launcherManager.GetConfig())
	stdlog.Printf("|*** 加载：HTTP服务：%s\n", httpConfig.GetAddr())

	// loggerForMiddleware
	loggerForMiddleware, err := launcherManager.GetLoggerForMiddleware()
	if err != nil {
		return nil, err
	}

	// authManager
	authManager, err := launcherManager.GetAuthManager()
	if err != nil {
		return nil, err
	}

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

	// setting
	settingConfig := configutil.SettingConfig(launcherManager.GetConfig())
	if settingConfig.GetEnableAuthMiddleware() {
		stdlog.Println("|*** 加载：验证中间件：HTTP")
		jwtMiddleware, err := middlewareutil.NewAuthMiddleware(authManager, authWhiteList)
		if err != nil {
			return srv, err
		}
		middlewareSlice = append(middlewareSlice, jwtMiddleware)
	}

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
