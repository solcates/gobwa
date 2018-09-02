package messages

import "testing"

func TestConfigResponse_Parse(t *testing.T) {
	type args struct {
		bin []byte
	}
	tests := []struct {
		name    string
		cr      *ConfigResponse
		args    args
		wantErr bool
	}{
		{
			name:    "ok",
			cr:      &ConfigResponse{},
			args:    args{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.cr.Parse(tt.args.bin); (err != nil) != tt.wantErr {
				t.Errorf("ConfigResponse.Parse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
