package configutil

import (
	configpb "github.com/go-micro-saas/service-kit/api/config"
	"testing"
)

// go test -v -count=1 ./config/ -test.run=TestLoadingConfigFromConsul
func TestLoadingConfigFromConsul(t *testing.T) {
	tests := []struct {
		name    string
		want    *configpb.Bootstrap
		wantErr bool
	}{
		{
			name:    "#loadingForConsul",
			want:    nil,
			wantErr: false,
		},
	}

	cfg := &configpb.Consul{
		Enable:             true,
		Address:            "127.0.0.1:8500",
		InsecureSkipVerify: true,
	}
	cc, err := newConsulClient(cfg)
	if err != nil {
		t.Fatal(err)
	}

	appConfig := &configpb.App{
		ConfigMethod:         CONFIG_METHOD_CONSUL,
		ConfigPathForGeneral: "go-micro-saas/general-config",
		ConfigPathForServer:  "go-micro-saas/ping-service/production/v1.0.0",
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadingConfigFromConsul(cc, appConfig)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadingConfigFromConsul() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("LoadingConfigFromConsul() got = %v, want %v", got, tt.want)
			//}
			if got.GetApp().GetServerName() == "" {
				t.Fatal("==> got.GetApp().GetServerName() is empty")
			}
			t.Log("==> got.GetApp().GetServerName(): ", got.GetApp().GetServerName())
		})
	}
}
