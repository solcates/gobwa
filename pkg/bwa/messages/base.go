package messages

//Message represents basic Messages that can be parse and serialized into the BWA protocol
type Message interface {
	//Parse parses the []byte slice into this struct
	Parse(in []byte) (err error)
	//Serialize serializes the struct into a []byte slice
	Serialize() (out []byte, err error)
}
