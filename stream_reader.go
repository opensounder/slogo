package slogo

import (
	"fmt"
	"io"
	"log"
)

//Decoder is the general interface
type Decoder interface {
	Decode() (*Frame, error)
}

type slDecoder struct {
	r         io.Reader
	version   int
	blocksize int
	cursor    int
}

//NewDecoder creates a new Decoder instance
func NewDecoder(r io.Reader, version, blocksize int) Decoder {
	return &slDecoder{r: r, version: version, blocksize: blocksize}
}

//Decode implements the Decode method
func (d *slDecoder) Decode() (*Frame, error) {
	p := make([]byte, d.blocksize)
	n, err := d.r.Read(p)
	if err != nil {
		if err == io.EOF {
			//todo: should perhaps handle any remainding bytes.
			fmt.Println(string(p[:n]))
			return nil, err
		}
		log.Panic(err)
	}
	fmt.Println("package")
	//todo: acutal frame reading
	frame := Frame{}
	frame.offset = -1

	return &frame, nil
}
