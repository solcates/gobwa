package messages

//ConfigResponse represents  the configuration response
type ConfigResponse struct {
}

//Parse parses the []byte slice into this struct
func (cr *ConfigResponse) Parse(bin []byte) (err error) {
	return
}

//Serialize serializes the struct into a []byte slice
func (cr *ConfigResponse) Serialize() (out []byte, err error) {
	return
}
