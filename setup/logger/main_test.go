package loggerutil

import (
	"os"
	"testing"
	"time"

	configpb "github.com/go-micro-saas/service-kit/api/config"
	"google.golang.org/protobuf/types/known/durationpb"
)

var (
	appConfig = &configpb.App{
		ProjectName:   "go-micro-saas",
		ServerName:    "test-service",
		ServerEnv:     "develop",
		ServerVersion: "v1.0.0",
		HttpEndpoints: nil,
		GrpcEndpoints: nil,
		Metadata:      nil,
	}
	logConfig = &configpb.Log{
		Console: &configpb.Log_Console{
			Enable: true,
			Level:  "DEBUG",
		},
		File: &configpb.Log_File{
			Enable:         true,
			Level:          "DEBUG",
			Dir:            "./runtime/logs",
			Filename:       "test",
			RotateTime:     durationpb.New(time.Hour * 24),
			RotateSize:     52428800,
			StorageAge:     durationpb.New(time.Hour * 24 * 30),
			StorageCounter: 10086,
		},
	}
	handler LoggerManager
)

func TestMain(m *testing.M) {
	var err error
	handler, err = NewLoggerManager(appConfig, logConfig)
	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}
