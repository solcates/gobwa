package bwa

import (
	"net"
	"reflect"
	"testing"
	"time"
)

func TestNewBalbowClient(t *testing.T) {
	type args struct {
		host   string
		port   int
		cancel chan bool
	}
	c := make(chan bool)
	tests := []struct {
		name string
		args args
		want *BalbowClient
	}{
		{
			name: "ok",
			args: args{
				host:   "172.16.1.21",
				port:   4257,
				cancel: c,
			},
			want: &BalbowClient{
				host:   "172.16.1.21",
				port:   4257,
				cancel: c,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewBalbowClient(tt.args.host, tt.args.port)
			got.cancel = tt.args.cancel
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBalbowClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBalbowClient_Connect(t *testing.T) {
	// Setup a Mock Server

	tests := []struct {
		name    string
		bc      *BalbowClient
		wantErr bool
	}{
		{
			name:    "ok",
			bc:      NewBalbowClient("172.16.1.21", 4257),
			wantErr: false,
		},
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
	bc.conn = &MockServer{}
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
	bc := NewBalbowClient("172.16.1.21", 4257)
	bc.conn = &MockServer{}
	tests := []struct {
		name string
		bc   *BalbowClient
	}{
		// TODO: Add test cases.
		{
			name: "ok",
			bc:   bc,
		},
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
	bc := NewBalbowClient("172.16.1.21", 4257)
	bc.conn = &MockServer{}
	tests := []struct {
		name string
		bc   *BalbowClient
	}{
		{
			name: "ok",
			bc:   bc,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.bc.RequestConfig()
		})
	}
}

func Test_chr(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := chr(tt.args.in); gotOut != tt.wantOut {
				t.Errorf("chr() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func TestBalbowClient_RequestControlInfo(t *testing.T) {
	bc := NewBalbowClient("172.16.1.21", 4257)
	bc.conn = &MockServer{}
	tests := []struct {
		name string
		bc   *BalbowClient
	}{
		{
			name: "ok",
			bc:   bc,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.bc.RequestControlInfo()
		})
	}
}

func TestBalbowClient_ToggleLight(t *testing.T) {
	tests := []struct {
		name string
		bc   *BalbowClient
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.bc.ToggleLight()
		})
	}
}

func TestBalbowClient_ToggleItem(t *testing.T) {
	type args struct {
		item string
	}
	tests := []struct {
		name string
		bc   *BalbowClient
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.bc.ToggleItem(tt.args.item)
		})
	}
}

type MockServer struct {
	data []byte
}

func (ms *MockServer) Read(b []byte) (n int, err error) {
	b = append(b, ms.data...)
	return
}

func (ms *MockServer) Write(b []byte) (n int, err error) {
	ms.data = b
	return
}

func (ms *MockServer) Close() error {
	panic("implement me")
}

func (ms *MockServer) LocalAddr() net.Addr {
	panic("implement me")
}

func (ms *MockServer) RemoteAddr() net.Addr {
	panic("implement me")
}

func (ms *MockServer) SetDeadline(t time.Time) error {
	panic("implement me")
}

func (ms *MockServer) SetReadDeadline(t time.Time) error {
	panic("implement me")
}

func (ms *MockServer) SetWriteDeadline(t time.Time) error {
	panic("implement me")
}

func TestBalbowClient_RequestControlConfig(t *testing.T) {
	tests := []struct {
		name string
		bc   *BalbowClient
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.bc.RequestControlConfig()
		})
	}
}
