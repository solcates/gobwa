package bwa

var defaultPort = 4257

type Client interface {
	Connect() (err error)
	Close() (err error)
}

type BalbowClient struct {
	host string
	port int
}

func NewBalbowClient(host string, port int) *BalbowClient {
	bc := &BalbowClient{
		host: host,
		port: port,
	}
	if bc.port == 0 {
		bc.port = defaultPort
	}
	return bc
}

func (bc *BalbowClient) Connect() (err error) {
	panic("implement me")
}

func (bc *BalbowClient) Close() (err error) {
	panic("implement me")
}
