package messages

//FilterCycleResponse is a response for the filtercycle commands
type FilterCycleResponse struct {
}

//Parse parses the []byte slice into this struct
func (fcr *FilterCycleResponse) Parse(bin []byte) (err error) {
	panic("implement me")
}

//Serialize serializes the struct into a []byte slice
func (fcr *FilterCycleResponse) Serialize() (out []byte, err error) {
	return
}
