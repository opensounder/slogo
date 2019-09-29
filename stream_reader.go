package slogo

import "io"

//Decoder is the general interface
type Decoder interface {
	Decode() ([]byte, error)
}

type slDecoder struct {
	r      io.Reader
	header Header
	cursor int
}

//NewDecoder creates a new Decoder instance
func NewDecoder(r io.Reader, header Header) Decoder {
	return &slDecoder{r: r, header: header}
}

func (d *slDecoder) Read(p []byte) (int, error) {
	n, err := d.r.Read(p)
	if err != nil {
		return n, err
	}
	return n, nil
}

//Decode implements the Decode method
func (d *slDecoder) Decode() ([]byte, error) {
	return nil, nil
}
