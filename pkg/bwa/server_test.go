package bwa

import (
	"reflect"
	"testing"
)

func TestNewBalboaServer(t *testing.T) {
	tests := []struct {
		name string
		want *BalboaServer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBalboaServer(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBalboaServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

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

func TestBalboaServer_SendMessage(t *testing.T) {
	tests := []struct {
		name string
		bs   *BalboaServer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.bs.SendMessage()
		})
	}
}
