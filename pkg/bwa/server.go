package bwa

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net"
)

//Server implements a basic API server
type Server interface {
	Run() (err error)
	Close()
}

//NewBalboaServer returns a default balboa server for later usage.
func NewBalboaServer(host string, port int) *BalboaServer {
	if host == "" {
		host = "127.0.0.1"
	}
	if port == 0 {
		port = 4257
	}
	bs := &BalboaServer{
		host: host,
		port: port,
	}
	return bs
}

//BalboaServer is the balboa server side implementation of the API
type BalboaServer struct {
	conn net.Conn
	port int
	host string
}

//Run will start the server and process inbound requests
func (bs *BalboaServer) Run() (err error) {

	//TODO: mock up server here
	var ln net.Listener
	add := fmt.Sprintf("%s:%d", bs.host, bs.port)
	ln, err = net.Listen("tcp", add)
	logrus.Infof("Listening on %s", add)
	for {
		bs.conn, err = ln.Accept()
		if err != nil {
			return
		}
		logrus.Infof("Server: Connection from %v", bs.conn.RemoteAddr().String())
		// handle connection request
		go bs.handlerequest()
	}
	return
}

//Close closes up the connection
func (bs *BalboaServer) Close() {
	bs.conn.Close()

}

func (bs *BalboaServer) handlerequest() {
	buf := []byte{}
	_, err := bs.conn.Read(buf)
	if err != nil {
		logrus.Error(err)
	}

}
