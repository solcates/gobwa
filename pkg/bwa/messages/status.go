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
	currentTemp    uint8
	priming        bool
	heatingMode    uint8
	tempScale      bool
	twentyFourHour bool
	heating        bool
	highRange      bool
	pump1          uint8
	pump2          uint8
	cp             bool
	light          bool
	hours          uint8
	minutes        uint8
	setTemp        uint8
}

//Parse the inbound array of bytes for it's status updates
func (s *Status) Parse(bin []byte) (err error) {
	if bin[1] != 0 {
		s.priming = true
	}
	//Currenttemp is the
	s.currentTemp = bin[2]

	// hours and minutes
	s.hours = bin[3]
	s.minutes = bin[4]

	// Heating mode
	switch bin[5] {
	case 0:
		s.heatingMode = Ready
	case 1:
		s.heatingMode = Rest
	case 2:
		s.heatingMode = ReadyInRest
	default:
		return errors.New("unknown heating mode")
	}

	// Statuses
	s.tempScale = (bin[9]&0x01 == 0x01)
	s.twentyFourHour = (bin[9]&0x02 == 0x02)
	s.heating = (bin[10]&0x30 != 0)
	s.highRange = (bin[10]&0x04 == 0x04)
	s.pump1 = bin[11] & 0x03
	s.pump2 = (bin[11] / 4) & 0x03
	s.cp = (bin[13]&0x02 == 0x02)
	s.light = (bin[14]&0x03 == 0x03)
	s.setTemp = bin[20]

	// if Celsius do the divide
	if s.tempScale {
		s.currentTemp = s.currentTemp / 2.0
		s.setTemp = s.setTemp / 2.0
	}
	return
}
