package messages

//ControlInfoResponse is expected after a successful control info request
type ControlInfoResponse struct {
}

//Parse will populate the struct with known mappings
func (cr *ControlInfoResponse) Parse(bin []byte) (err error) {
	return
}
