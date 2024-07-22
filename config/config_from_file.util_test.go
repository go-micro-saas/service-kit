package configutil

import (
	"testing"

	configpb "github.com/go-micro-saas/service-kit/api/config"
)

// go test -v -count=1 ./config/ -test.run=TestLoadingFile
func TestLoadingFile(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    *configpb.Bootstrap
		wantErr bool
	}{
		{
			name: "#TestLoadingFile",
			args: args{
				filePath: "config_example.yaml",
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadingFile(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadingFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("LoadingFile() got = %v, want %v", got, tt.want)
			//}
			t.Logf("Boostrap.App: %#v\n", got.App)
		})
	}
}
