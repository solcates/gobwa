package messages

import (
	"errors"
	"github.com/sirupsen/logrus"
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
	tempScale      uint8
	twentyFourHour uint8
	heating        uint8
	tempRange      uint8
	pump1          uint8
	pump2          uint8
	cp             bool
	light          bool
	hours          uint8
	minutes        uint8
}

//Parse the inbound array of bytes for it's status updates
func (s *Status) Parse(bin []byte) (err error) {
	if bin[6] != 0 {
		s.priming = true
	}
	//Currenttempt
	s.currentTemp = bin[7]

	// hours and minutes
	s.hours = bin[8]
	s.minutes = bin[9]

	// heating mode
	switch bin[10] {
	case 0:
		s.heatingMode = Ready
	case 1:
		s.heatingMode = Rest
	case 2:
		s.heatingMode = ReadyInRest
	default:
		return errors.New("unknown heating mode")
	}

	// F vs C
	s.tempScale = bin[14]

	//

	logrus.Infof("currentTemp: %v", s.currentTemp)
	logrus.Infof("x: %v", s.tempScale)
	logrus.Info("")
	return
}
