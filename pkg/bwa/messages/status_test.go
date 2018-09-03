package messages

import (
	"reflect"
	"testing"
)

func TestStatus_Parse(t *testing.T) {
	type args struct {
		bin []byte
	}
	tests := []struct {
		name    string
		s       *Status
		args    args
		wantErr bool
		want    *Status
	}{
		{
			name: "ok",
			s:    &Status{},
			args: args{
				bin: []byte{

					0x00, 0x00, // ??
					0x67, // Current Temp
					0x0a, 0x2f, 0x00, 0x00, 0x01, 0x00, 0x00,
					0x0c, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x08, 0x66, 0x00, 0x00, 0x10,
				},
			},
			wantErr: false,
			want: &Status{
				CurrentTemp:    103,
				Priming:        false,
				HeatingMode:    0,
				TempScale:      false,
				TwentyFourHour: false,
				Heating:        false,
				HighRange:      true,
				Pump1:          1,
				Pump2:          0,
				Cp:             false,
				Light:          false,
				Hours:          10,
				Minutes:        47,
				SetTemp:        102,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Parse(tt.args.bin); (err != nil) != tt.wantErr {
				t.Errorf("Status.Parse() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.s, tt.want) {
				t.Errorf("Parse() = %v, want %v", tt.s, tt.want)
			}
		})
	}
}
