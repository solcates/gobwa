package bwa

import "net"

type Server interface {
	Run() (err error)
	SendMessage()
	Close()
}

func NewBalboaServer() *BalboaServer {
	bs := &BalboaServer{}
	return bs
}

type BalboaServer struct {
	conn net.Conn
}

func (bs *BalboaServer) Run() (err error) {

	//TODO: mock up server here
	var ln net.Listener
	ln, err = net.Listen("tcp", ":4257")
	for {
		bs.conn, err = ln.Accept()
		if err != nil {
			return
		}
		// handle connection request
		go bs.handlerequest()
	}
	return
}
func (bs *BalboaServer) Close() {

}

func (bs *BalboaServer) handlerequest() {

}

func (bs *BalboaServer) SendMessage() {

	panic("implement me")
}
