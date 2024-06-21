package configutil

import (
	"testing"
)

// go test -v -count=1 ./config/ -test.run=TestSetupWithFile
func TestSetupWithFile(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    *Bootstrap
		wantErr bool
	}{
		{
			name: "#TestSetupWithFile",
			args: args{
				filePath: "./config",
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SetupWithFile(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetupWithFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("SetupWithFile() got = %v, want %v", got, tt.want)
			//}
			t.Logf("Boostrap.App: %#v\n", got.App)
		})
	}
}
