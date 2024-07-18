package configutil

import (
	configpb "github.com/go-micro-saas/service-kit/api/config"
	apppkg "github.com/ikaiguang/go-srv-kit/kratos/app"
)

type ConfigManager interface {
	Env() apppkg.RuntimeEnvEnum_RuntimeEnv
	IsLocalMode() bool
	IsDebugMode() bool

	AppConfig() *configpb.App
	SettingConfig() *configpb.Setting
	SettingCaptchaConfig() *configpb.Setting_Captcha
	SettingLoginConfig() *configpb.Setting_Login

	HTTPConfig() *configpb.Server_HTTP
	GRPCConfig() *configpb.Server_GRPC

	LogConfig() *configpb.Log
	LogConsoleConfig() *configpb.Log_Console
	LogFileConfig() *configpb.Log_File

	MysqlConfig() *configpb.MySQL
	PostgresConfig() *configpb.PSQL
	RedisConfig() *configpb.Redis
	RabbitMQConfig() *configpb.Rabbitmq
	ConsulConfig() *configpb.Consul
	EtcdConfig() *configpb.Etcd
	Jaeger() *configpb.Jaeger

	TransferEncryptConfig() *configpb.Encrypt_TransferEncrypt
	ServiceEncryptConfig() *configpb.Encrypt_ServiceEncrypt
	TokenEncryptConfig() *configpb.Encrypt_TokenEncrypt

	ClusterServiceEndpoints() []*configpb.ClientApi_Endpoint
	ThirdPartyEndpoints() []*configpb.ClientApi_Endpoint
	SnowflakeConfig() *configpb.Snowflake
}
