package configutil

import (
	_ "embed"
	"os"
	"testing"
)

var (
	//go:embed config_example.yaml
	configBuf  []byte
	configPath string = "config_example.yaml"
	handler    ConfigManager
)

func TestMain(m *testing.M) {
	boostrap, err := LoadingFile(configPath)
	if err != nil {
		panic(err)
	}

	SetBootstrap(boostrap)
	handler = NewConfigManager(boostrap)

	os.Exit(m.Run())
}
