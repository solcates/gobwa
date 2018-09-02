package bwa

import (
	"fmt"
	"github.com/gdamore/encoding"
	"github.com/sigurn/crc8"
	"github.com/sirupsen/logrus"
	"log"
	"net"
)

var defaultPort = 4257
var leftoverData = ""

//Client interface for Balboa
type Client interface {
	Connect() (err error)
	SendMessage(message string) (err error)
	Close() (err error)
}

//BalbowClient implements a TCP Socket to port 4257 on a Balboa device.
type BalbowClient struct {
	host string
	port int
	conn net.Conn
}

//NewBalbowClient returns a default client
func NewBalbowClient(host string, port int) *BalbowClient {
	bc := &BalbowClient{
		host: host,
		port: port,
	}
	if bc.port == 0 {
		bc.port = defaultPort
	}
	if bc.host == "" {
		bc.host = "127.0.0.1"
	}
	return bc
}

//Connect establishes a socket to the port and host of this client
func (bc *BalbowClient) Connect() (err error) {
	service := fmt.Sprintf("%v:%v", bc.host, bc.port)
	//log.Println(service)
	var tcpAddr *net.TCPAddr
	tcpAddr, err = net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		return
	}
	bc.conn, err = net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return
	}

	log.Println("Connected")
	// start polling...
	go bc.poll()

	return
}

func (bc *BalbowClient) poll() {
	for {
		buf := make([]byte, 128)
		_, err := bc.conn.Read(buf)
		if err != nil {
			return
		}
		//fmt.Println("N: %v", n)
		//fmt.Println("read: % x", buf)
		Parse(buf)
	}
	return
}

//SendMessage serializes and writes the message to the client's open socket.
// It will also estabolish the connection again if not connected.
func (bc *BalbowClient) SendMessage(message string) (err error) {
	if bc.conn == nil {
		err = bc.Connect()
		if err != nil {
			return
		}
	}
	fMessage := prepMessage(message)
	_, err = bc.conn.Write([]byte(fMessage))
	if err != nil {
		return
	}

	// Read Response as a channel call back
	//result, err := ioutil.ReadAll(bc.conn)
	//if err != nil {
	//	return
	//}
	//log.Println(result)
	return
}

func prepMessage(message string) string {

	// The message length will need 2 more added for the start and end bits
	length := len(message) + 2
	t := chr(string(length))
	//
	fullMessage := fmt.Sprintf("%s%s", t, message)

	// CRC8 the full message with Init and XorMask of 0x02
	params := crc8.CRC8
	params.Init = 0x02
	params.XorOut = 0x02
	table := crc8.MakeTable(params)

	checksum := crc8.Checksum([]byte(fullMessage), table)
	checksumChar := chr(string(checksum))
	finalMessage := fmt.Sprintf("%s%s%s%s", start, fullMessage, checksumChar, end)
	return finalMessage
}

func chr(in string) (out string) {
	out, _ = encoding.ASCII.NewEncoder().String(in)
	return
}

//Close closes the client
func (bc *BalbowClient) Close() (err error) {
	return bc.conn.Close()
}

//RequestConfig resquests the config from the SPA.
//
// if the configuraiton i
func (bc *BalbowClient) RequestConfig() {
	bc.SendMessage("\x0a\xbf\x04")
}

//RequestControlInfo reqeustes the Control INfo fo rthe maction
func (bc *BalbowClient) RequestControlInfo() {
	logrus.Debug("RCI")
	bc.SendMessage("\x0a\xbf\x22\x02\x00\x00")
}

func (bc *BalbowClient) ToggleLight() {
	logrus.Debug("ToggleLight")
	bc.ToggleItem("\x11\x00")
}

func (bc *BalbowClient) ToggleItem(item string) {
	bc.SendMessage("x0a\xbf\x11" + item)
}
