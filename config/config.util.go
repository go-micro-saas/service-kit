package configutil

import (
	"sync"

	configpb "github.com/go-micro-saas/service-kit/api/config"
	apppkg "github.com/ikaiguang/go-srv-kit/kratos/app"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
)

var (
	_bootstrap *configpb.Bootstrap

	// 不要直接使用 s.env, 请使用 Env()
	_env     apppkg.RuntimeEnvEnum_RuntimeEnv
	_envOnce sync.Once
)

func SetBootstrap(bootstrap *configpb.Bootstrap) {
	_bootstrap = bootstrap
}

func GetBootstrap() (*configpb.Bootstrap, error) {
	if _bootstrap == nil {
		e := errorpkg.ErrorUninitialized("bootstrap is uninitialized")
		return nil, errorpkg.WithStack(e)
	}
	return getBootstrap(), nil
}

func getBootstrap() *configpb.Bootstrap {
	return _bootstrap
}

func Env() apppkg.RuntimeEnvEnum_RuntimeEnv {
	_envOnce.Do(func() {
		_env = apppkg.ParseEnv(getBootstrap().GetApp().GetServerEnv())
	})
	return _env
}

func IsDebugMode() bool {
	switch Env() {
	default:
		return false
	case apppkg.RuntimeEnvEnum_LOCAL, apppkg.RuntimeEnvEnum_DEVELOP, apppkg.RuntimeEnvEnum_TESTING:
		return true
	}
}

func IsLocalMode() bool {
	switch Env() {
	default:
		return false
	case apppkg.RuntimeEnvEnum_LOCAL:
		return true
	}
}

func AppConfig() *configpb.App {
	return getBootstrap().GetApp()
}

func SettingConfig() *configpb.Setting {
	return getBootstrap().GetSetting()
}
func SettingCaptchaConfig() *configpb.Setting_Captcha {
	return getBootstrap().GetSetting().GetCaptcha()
}
func SettingLoginConfig() *configpb.Setting_Login {
	return getBootstrap().GetSetting().GetLogin()
}

func HTTPConfig() *configpb.Server_HTTP {
	return getBootstrap().GetServer().GetHttp()
}
func GRPCConfig() *configpb.Server_GRPC {
	return getBootstrap().GetServer().GetGrpc()
}

func LogConfig() *configpb.Log {
	return getBootstrap().GetLog()
}
func LogConsoleConfig() *configpb.Log_Console {
	return getBootstrap().GetLog().GetConsole()
}
func LogFileConfig() *configpb.Log_File {
	return getBootstrap().GetLog().GetFile()
}

func MysqlConfig() *configpb.MySQL {
	return getBootstrap().GetMysql()
}
func PostgresConfig() *configpb.PSQL {
	return getBootstrap().GetPsql()
}
func RedisConfig() *configpb.Redis {
	return getBootstrap().GetRedis()
}
func RabbitMQConfig() *configpb.Rabbitmq {
	return getBootstrap().GetRabbitmq()
}
func ConsulConfig() *configpb.Consul {
	return getBootstrap().GetConsul()
}
func EtcdConfig() *configpb.Etcd {
	return getBootstrap().GetEtcd()
}
func Jaeger() *configpb.Jaeger {
	return getBootstrap().GetJaeger()
}

func TransferEncryptConfig() *configpb.Encrypt_TransferEncrypt {
	return getBootstrap().GetEncrypt().GetTransferEncrypt()
}
func ServiceEncryptConfig() *configpb.Encrypt_ServiceEncrypt {
	return getBootstrap().GetEncrypt().GetServiceEncrypt()
}
func TokenEncryptConfig() *configpb.Encrypt_TokenEncrypt {
	return getBootstrap().GetEncrypt().GetTokenEncrypt()
}

func ClusterServiceEndpoints() []*configpb.ClientApi_Endpoint {
	return getBootstrap().GetClientApi().ClusterService
}
func ThirdPartyEndpoints() []*configpb.ClientApi_Endpoint {
	return getBootstrap().GetClientApi().ThirdParty
}

func SnowflakeConfig() *configpb.Snowflake {
	return getBootstrap().GetSnowflake()
}
