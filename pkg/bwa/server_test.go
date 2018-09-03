package bwa

import (
	"reflect"
	"testing"
)

func TestBalboaServer_Run(t *testing.T) {
	tests := []struct {
		name    string
		bs      *BalboaServer
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.bs.Run(); (err != nil) != tt.wantErr {
				t.Errorf("BalboaServer.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBalboaServer_Close(t *testing.T) {
	tests := []struct {
		name string
		bs   *BalboaServer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.bs.Close()
		})
	}
}

func TestBalboaServer_handlerequest(t *testing.T) {

	tests := []struct {
		name string
		bs   *BalboaServer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.bs.handlerequest()
		})
	}
}

func TestNewBalboaServer(t *testing.T) {
	type args struct {
		host string
		port int
	}
	tests := []struct {
		name string
		args args
		want *BalboaServer
	}{
		{
			name: "Default",
			args: args{},
			want: &BalboaServer{
				host: "127.0.0.1",
				port: 4257,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBalboaServer(tt.args.host, tt.args.port); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBalboaServer() = %v, want %v", got, tt.want)
			}
		})
	}
}
