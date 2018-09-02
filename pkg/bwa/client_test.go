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
				host: "172.16.1.21",
				port: 4257,
			},
			want: &BalbowClient{
				host: "172.16.1.21",
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
	// Setup a Mock Server
	var bs *BalboaServer
	go func() {
		bs = NewBalboaServer()
		bs.Run()
	}()

	defer bs.Close()

	tests := []struct {
		name    string
		bc      *BalbowClient
		wantErr bool
	}{
		//{
		//	name:    "ok",
		//	bc:      NewBalbowClient("172.16.1.21", 4257),
		//	wantErr: false,
		//},
	}
	// run server in background
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			//

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

func TestBalbowClient_SendMessage(t *testing.T) {
	bc := NewBalbowClient("172.16.1.21", 4257)
	tests := []struct {
		name    string
		message string
		bc      *BalbowClient
		wantErr bool
	}{
		{
			name:    "ok",
			bc:      bc,
			wantErr: false,
			message: "\x0a\xbf\x04",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.bc.SendMessage(tt.message); (err != nil) != tt.wantErr {
				t.Errorf("BalbowClient.SendMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBalbowClient_poll(t *testing.T) {
	tests := []struct {
		name string
		bc   *BalbowClient
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.bc.poll()
		})
	}
}

func Test_prepMessage(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "RequestConfig",
			args: args{
				message: "\x0a\xbf\x04",
			},
			want: "\x7e\x05\x0a\xbf\x04\x77\x7e",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prepMessage(tt.args.message); got != tt.want {
				t.Errorf("prepMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBalbowClient_RequestConfig(t *testing.T) {
	tests := []struct {
		name string
		bc   *BalbowClient
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.bc.RequestConfig()
		})
	}
}
