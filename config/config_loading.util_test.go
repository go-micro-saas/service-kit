package configutil

import (
	"os"
	"testing"
)

// go test -v -count=1 ./config/ -test.run=TestCurrentPath
func TestCurrentPath(t *testing.T) {
	// get $GOPATH
	gopath := os.Getenv("GOPATH")
	// get $GOPATH/src/github.com/go-micro-saas/service-kit/config
	tests := []struct {
		name string
		want string
	}{
		{
			name: "#TestCurrentPath",
			want: gopath + "/src/github.com/go-micro-saas/service-kit/config",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CurrentPath(); got != tt.want {
				t.Errorf("CurrentPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
