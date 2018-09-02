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
		// TODO: Add test cases.
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
