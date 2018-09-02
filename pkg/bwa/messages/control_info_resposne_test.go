package messages

import "testing"

func TestControlInfoResponse_Parse(t *testing.T) {
	type args struct {
		bin []byte
	}
	tests := []struct {
		name    string
		cr      *ControlInfoResponse
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.cr.Parse(tt.args.bin); (err != nil) != tt.wantErr {
				t.Errorf("ControlInfoResponse.Parse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
