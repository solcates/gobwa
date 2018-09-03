package messages

//ControlInfoResponse is expected after a successful control info request
type ControlInfoResponse struct {
}

//Parse parses the []byte slice into this struct
func (cr *ControlInfoResponse) Parse(bin []byte) (err error) {
	return
}

//Serialize serializes the struct into a []byte slice
func (cr *ControlInfoResponse) Serialize() (out []byte, err error) {
	return
}
