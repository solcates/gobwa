package messages

type Message interface {
	Parse(bin []byte) (err error)
}
