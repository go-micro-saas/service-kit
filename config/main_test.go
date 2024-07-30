package configutil

import (
	_ "embed"
	stdlog "log"
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
		stdlog.Printf("%+v\n", err)
		panic(err)
	}

	SetConfig(boostrap)

	os.Exit(m.Run())
}
