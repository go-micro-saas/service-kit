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

func SetConfig(bootstrap *configpb.Bootstrap) {
	_bootstrap = bootstrap
}

func GetConfig() (*configpb.Bootstrap, error) {
	if _bootstrap == nil {
		e := errorpkg.ErrorUninitialized("bootstrap is uninitialized")
		return nil, errorpkg.WithStack(e)
	}
	return getConfig(), nil
}

func getConfig() *configpb.Bootstrap {
	return _bootstrap
}

func Env() apppkg.RuntimeEnvEnum_RuntimeEnv {
	_envOnce.Do(func() {
		_env = apppkg.ParseEnv(getConfig().GetApp().GetServerEnv())
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
	return getConfig().GetApp()
}

func SettingConfig() *configpb.Setting {
	return getConfig().GetSetting()
}
func SettingCaptchaConfig() *configpb.Setting_Captcha {
	return getConfig().GetSetting().GetCaptcha()
}
func SettingLoginConfig() *configpb.Setting_Login {
	return getConfig().GetSetting().GetLogin()
}

func HTTPConfig() *configpb.Server_HTTP {
	return getConfig().GetServer().GetHttp()
}
func GRPCConfig() *configpb.Server_GRPC {
	return getConfig().GetServer().GetGrpc()
}

func LogConfig() *configpb.Log {
	return getConfig().GetLog()
}
func LogConsoleConfig() *configpb.Log_Console {
	return getConfig().GetLog().GetConsole()
}
func LogFileConfig() *configpb.Log_File {
	return getConfig().GetLog().GetFile()
}

func MysqlConfig() *configpb.MySQL {
	return getConfig().GetMysql()
}
func PostgresConfig() *configpb.PSQL {
	return getConfig().GetPsql()
}
func RedisConfig() *configpb.Redis {
	return getConfig().GetRedis()
}
func RabbitmqConfig() *configpb.Rabbitmq {
	return getConfig().GetRabbitmq()
}
func ConsulConfig() *configpb.Consul {
	return getConfig().GetConsul()
}
func EtcdConfig() *configpb.Etcd {
	return getConfig().GetEtcd()
}
func JaegerConfig() *configpb.Jaeger {
	return getConfig().GetJaeger()
}

func TransferEncryptConfig() *configpb.Encrypt_TransferEncrypt {
	return getConfig().GetEncrypt().GetTransferEncrypt()
}
func ServiceEncryptConfig() *configpb.Encrypt_ServiceEncrypt {
	return getConfig().GetEncrypt().GetServiceEncrypt()
}
func TokenEncryptConfig() *configpb.Encrypt_TokenEncrypt {
	return getConfig().GetEncrypt().GetTokenEncrypt()
}

func ClusterClientApis() []*configpb.ClusterClientApi {
	return getConfig().GetClusterClientApi()
}
func ThirdPartyApis() []*configpb.ThirdPartyApi {
	return getConfig().GetThirdPartyApi()
}

func SnowflakeConfig() *configpb.Snowflake {
	return getConfig().GetSnowflake()
}
