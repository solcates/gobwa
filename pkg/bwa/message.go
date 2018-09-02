package bwa

import (
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/solcates/gobwa/pkg/bwa/messages"
)

const (
	start = "\x7e"
	end   = "\x7e"
)

//Message

//Parse inspects the incoming byte slice for a known message type and it's contents
func Parse(bin []byte) (message messages.Message, err error) {

	// Ensure the start and end are there in the message:
	if bin[0] != 0x7e {
		err = errors.New("bad interface")
		return
	}

	// message type is the 3 numbers of offset 2-4 (skip the first 2)
	//length := bin[1]

	messaegType := bin[2:5]

	//fmt.Printf("Length: %v\n", length)
	//fmt.Printf("Type: %v\n", messaegType)

	mt := string(messaegType)
	switch mt {
	case "\x0a\xbf\x94":
		logrus.Debug("Returning ConfigResponse")
		message = &messages.ConfigResponse{}

		return
	case "\xff\xaf\x13":
		logrus.Debug("Returning Status")
		message = &messages.Status{}
	default:
		logrus.Errorf("Unknown Message Type: %x", mt)
	}
	err = message.Parse(bin)

	return
}
