package configutil

import (
	_ "embed"
	"os"
	"testing"
)

var (
	//go:embed config_example.yaml
	configBuf  []byte
	configPath = "config_example.yaml"
)

func TestMain(m *testing.M) {
	boostrap, err := LoadingFile(configPath)
	if err != nil {
		panic(err)
	}

	SetBootstrap(boostrap)

	os.Exit(m.Run())
}
