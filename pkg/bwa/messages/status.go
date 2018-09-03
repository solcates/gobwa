package messages

import (
	"errors"
)

const (
	//Ready is when the spa is ready for swimmers
	Ready = iota
	//Rest is when the spa is saving energy
	Rest
	//ReadyInRest is hybrid of the 2
	ReadyInRest
)

//Status represents current status responses sent by the spa
type Status struct {
	CurrentTemp    uint8
	Priming        bool
	HeatingMode    uint8
	TempScale      bool
	TwentyFourHour bool
	Heating        bool
	HighRange      bool
	Pump1          uint8
	Pump2          uint8
	Cp             bool
	Light          bool
	Hours          uint8
	Minutes        uint8
	SetTemp        uint8
}

//Parse parses the []byte slice into this struct
func (s *Status) Parse(bin []byte) (err error) {
	if bin[1] != 0 {
		s.Priming = true
	}
	//Currenttemp is the
	s.CurrentTemp = bin[2]

	// hours and minutes
	s.Hours = bin[3]
	s.Minutes = bin[4]

	// Heating mode
	switch bin[5] {
	case 0:
		s.HeatingMode = Ready
	case 1:
		s.HeatingMode = Rest
	case 2:
		s.HeatingMode = ReadyInRest
	default:
		return errors.New("unknown heating mode")
	}

	// Statuses
	s.TempScale = (bin[9]&0x01 == 0x01)
	s.TwentyFourHour = (bin[9]&0x02 == 0x02)
	s.Heating = (bin[10]&0x30 != 0)
	s.HighRange = (bin[10]&0x04 == 0x04)
	s.Pump1 = bin[11] & 0x03
	s.Pump2 = (bin[11] / 4) & 0x03
	s.Cp = (bin[13]&0x02 == 0x02)
	s.Light = (bin[14]&0x03 == 0x03)
	s.SetTemp = bin[20]

	// if Celsius do the divide
	if s.TempScale {
		s.CurrentTemp = s.CurrentTemp / 2.0
		s.SetTemp = s.SetTemp / 2.0
	}
	return
}

//Serialize serialzes the struct to a []byte slice
func (s *Status) Serialize() (out []byte, err error) {
	return
}
