package bwa

import (
	"reflect"
	"testing"
)

func TestBalboaMessage_Parse(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name    string
		bm      *BalboaMessage
		args    args
		wantOut []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := tt.bm.Parse(tt.args.in); !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("BalboaMessage.Parse() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func TestBalboaMessage_Serialize(t *testing.T) {
	type args struct {
		in []byte
	}
	tests := []struct {
		name    string
		bm      *BalboaMessage
		args    args
		wantOut string
	}{
		{
			name:    "ok",
			bm:      &BalboaMessage{

			},
			args:    args{},
			wantOut: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := tt.bm.Serialize(tt.args.in); gotOut != tt.wantOut {
				t.Errorf("BalboaMessage.Serialize() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
