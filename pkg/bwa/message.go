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

	messageType := bin[2:5]
	messagePayload := bin[5 : len(bin)-2]
	mt := string(messageType)
	switch mt {

	case "\x0a\xbf\x94":
		logrus.Debug("Received ConfigResponse")
		message = &messages.ConfigResponse{}
	case "\xff\xaf\x13":
		logrus.Debug("Received Status")
		message = &messages.Status{}
	case "\x0a\xbf\x24":
		logrus.Debugf("Received ControlInfoResponse")
		message = &messages.ControlInfoResponse{}
	case "\x0a\xbf\x23":
		logrus.Debugf("Received FilterCycleResponse")
		message = &messages.FilterCycleResponse{}

	default:
		logrus.Errorf("Unknown Message Type: %x", mt)
	}
	err = message.Parse(messagePayload)

	return
}
