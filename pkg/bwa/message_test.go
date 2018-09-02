package bwa

import (
	"reflect"
	"testing"

	"github.com/solcates/gobwa/pkg/bwa/messages"
)

func TestParse(t *testing.T) {
	type args struct {
		bin []byte
	}
	tests := []struct {
		name        string
		args        args
		wantMessage messages.Message
		wantErr     bool
	}{
		{
			name: "Status",
			args: args{
				bin: []byte{
					0x7e,
					0x1a,
					0xff, 0xaf, 0x13,
					0x00, 0x00, // ??
					0x67, // Current Temp
					0x0a, 0x2f, 0x00, 0x00, 0x01, 0x00, 0x00,
					0x0c, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x08, 0x66, 0x00, 0x00, 0x10,
					0xcb, 0x7e,
				},
			},
			wantMessage: nil,
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMessage, err := Parse(tt.args.bin)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotMessage, tt.wantMessage) {
				t.Errorf("Parse() = %v, want %v", gotMessage, tt.wantMessage)
			}
		})
	}
}
