package setuputil

import (
	configutil "github.com/go-micro-saas/service-kit/config"
	"path/filepath"
	"testing"
)

// go test -v -count=1 ./setup/ -test.run=TestSetup
func TestSetup(t *testing.T) {
	confPath := configutil.CurrentPath()
	confPath = filepath.Join(confPath, "config_example.yaml")
	type args struct {
		configFilePath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "#testingSetup",
			args:    args{configFilePath: confPath},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Setup(tt.args.configFilePath); (err != nil) != tt.wantErr {
				t.Errorf("Setup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
