package bwa

import "github.com/solcates/broadcast"

type Discoverer struct {
	*broadcast.UDPBroadcaster
}

func NewDiscoverer() *Discoverer {
	bc := broadcast.NewUDPBroadcaster(30303, "Discovery: Who is out there?")
	return &Discoverer{
		UDPBroadcaster: bc,
	}
}
