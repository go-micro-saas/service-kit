package configutil

import (
	configpb "github.com/go-micro-saas/service-kit/api/config"
	apppkg "github.com/ikaiguang/go-srv-kit/kratos/app"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	"sync"
)

var (
	_bootstrap *configpb.Bootstrap
)

func SetBootstrap(bootstrap *configpb.Bootstrap) {
	_bootstrap = bootstrap
}

func GetBootstrap() (*configpb.Bootstrap, error) {
	if _bootstrap == nil {
		e := errorpkg.ErrorUninitialized("bootstrap is uninitialized")
		return nil, errorpkg.WithStack(e)
	}
	return _bootstrap, nil
}

// configManager 实现 Config
type configManager struct {
	conf *configpb.Bootstrap

	// 不要直接使用 s.env, 请使用 Env()
	env     apppkg.RuntimeEnvEnum_RuntimeEnv
	envOnce sync.Once
}

func NewConfigManager(conf *configpb.Bootstrap) ConfigManager {
	if conf == nil {
		panic("conf is nil")
	}
	return &configManager{conf: conf}
}

func (s *configManager) Env() apppkg.RuntimeEnvEnum_RuntimeEnv {
	s.envOnce.Do(func() {
		s.env = apppkg.ParseEnv(s.conf.GetApp().GetServerEnv())
	})
	return s.env
}

func (s *configManager) IsDebugMode() bool {
	switch s.Env() {
	default:
		return false
	case apppkg.RuntimeEnvEnum_LOCAL, apppkg.RuntimeEnvEnum_DEVELOP, apppkg.RuntimeEnvEnum_TESTING:
		return true
	}
}

func (s *configManager) IsLocalMode() bool {
	switch s.Env() {
	default:
		return false
	case apppkg.RuntimeEnvEnum_LOCAL:
		return true
	}
}

func (s *configManager) AppConfig() *configpb.App {
	return s.conf.GetApp()
}

func (s *configManager) SettingConfig() *configpb.Setting {
	return s.conf.GetSetting()
}
func (s *configManager) SettingCaptchaConfig() *configpb.Setting_Captcha {
	return s.conf.GetSetting().GetCaptcha()
}
func (s *configManager) SettingLoginConfig() *configpb.Setting_Login {
	return s.conf.GetSetting().GetLogin()
}

func (s *configManager) HTTPConfig() *configpb.Server_HTTP {
	return s.conf.GetServer().GetHttp()
}
func (s *configManager) GRPCConfig() *configpb.Server_GRPC {
	return s.conf.GetServer().GetGrpc()
}

func (s *configManager) LogConfig() *configpb.Log {
	return s.conf.GetLog()
}
func (s *configManager) LogConsoleConfig() *configpb.Log_Console {
	return s.conf.GetLog().GetConsole()
}
func (s *configManager) LogFileConfig() *configpb.Log_File {
	return s.conf.GetLog().GetFile()
}

func (s *configManager) MysqlConfig() *configpb.MySQL {
	return s.conf.GetMysql()
}
func (s *configManager) PostgresConfig() *configpb.PSQL {
	return s.conf.GetPsql()
}
func (s *configManager) RedisConfig() *configpb.Redis {
	return s.conf.GetRedis()
}
func (s *configManager) RabbitMQConfig() *configpb.Rabbitmq {
	return s.conf.GetRabbitmq()
}
func (s *configManager) ConsulConfig() *configpb.Consul {
	return s.conf.GetConsul()
}
func (s *configManager) EtcdConfig() *configpb.Etcd {
	return s.conf.GetEtcd()
}
func (s *configManager) Jaeger() *configpb.Jaeger {
	return s.conf.GetJaeger()
}

func (s *configManager) TransferEncryptConfig() *configpb.Encrypt_TransferEncrypt {
	return s.conf.GetEncrypt().GetTransferEncrypt()
}
func (s *configManager) ServiceEncryptConfig() *configpb.Encrypt_ServiceEncrypt {
	return s.conf.GetEncrypt().GetServiceEncrypt()
}
func (s *configManager) TokenEncryptConfig() *configpb.Encrypt_TokenEncrypt {
	return s.conf.GetEncrypt().GetTokenEncrypt()
}

func (s *configManager) ClusterServiceEndpoints() []*configpb.ClientApi_Endpoint {
	return s.conf.GetClientApi().ClusterService
}
func (s *configManager) ThirdPartyEndpoints() []*configpb.ClientApi_Endpoint {
	return s.conf.GetClientApi().ThirdParty
}

func (s *configManager) SnowflakeConfig() *configpb.Snowflake {
	return s.conf.GetSnowflake()
}
