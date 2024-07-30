package configutil

import (
	"testing"
)

// go test -v -count=1 ./config/ -test.run=Test_configManager_IsDebugMode
func Test_configManager_IsDebugMode(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{
			name: "#IsDebugModel",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsDebugMode(); got != tt.want {
				t.Errorf("IsDebugMode() = %v, want %v", got, tt.want)
			}
			t.Log("config env: ", Env())
			conf, err := GetConfig()
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			t.Log("==> GetClusterClientApi length: ", len(conf.GetClusterClientApi()))
			t.Log("==> GetThirdPartyApi length: ", len(conf.GetThirdPartyApi()))
			t.Log("==> conf.GetApp().Id: ", conf.GetApp().GetId())
			t.Log("==> conf.GetApp().GetMetadata: ", conf.GetApp().GetMetadata())
		})
	}
}
