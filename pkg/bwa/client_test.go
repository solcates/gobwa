package bwa

import (
	"reflect"
	"testing"
)

func TestNewBalbowClient(t *testing.T) {
	type args struct {
		host string
		port int
	}
	tests := []struct {
		name string
		args args
		want *BalbowClient
	}{
		{
			name: "ok",
			args: args{
				host: "",
				port: 4257,
			},
			want: &BalbowClient{
				host: "localhost",
				port: 4257,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBalbowClient(tt.args.host, tt.args.port); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBalbowClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBalbowClient_Connect(t *testing.T) {
	tests := []struct {
		name    string
		bc      *BalbowClient
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.bc.Connect(); (err != nil) != tt.wantErr {
				t.Errorf("BalbowClient.Connect() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBalbowClient_Close(t *testing.T) {
	tests := []struct {
		name    string
		bc      *BalbowClient
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.bc.Close(); (err != nil) != tt.wantErr {
				t.Errorf("BalbowClient.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}