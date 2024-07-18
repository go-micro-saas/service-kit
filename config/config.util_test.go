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
			if got := handler.IsDebugMode(); got != tt.want {
				t.Errorf("IsDebugMode() = %v, want %v", got, tt.want)
			}
			t.Log("config env: ", handler.Env())
		})
	}
}
