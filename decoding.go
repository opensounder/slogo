package slogo

import (
	"encoding/binary"
	"io"
	"os"
)

//Decoder is the general interface
type Decoder interface {
	Next(FrameReader) error
}

//OpenLog is a function
func OpenLog(path string) (logfile *os.File, header Header, err error) {
	logfile, err = os.Open(path)
	header = Header{}
	if err != nil {
		return
	}

	header, err = ReadHeader(logfile)
	if err != nil {
		logfile.Close()
		logfile = nil
		return
	}

	return
}

//ReadHeader is another function
func ReadHeader(r io.Reader) (header Header, err error) {
	header = Header{}
	err = binary.Read(r, binary.LittleEndian, &header)
	return
}

type slDecoder struct {
	r      io.Reader
	header *Header
}

func OpenDecoder(filename string) (*os.File, Decoder, error) {
	logfile, header, err := OpenLog(filename)
	if err != nil {
		return logfile, nil, err
	}
	d := NewDecoder(logfile, header)
	return logfile, d, nil
}

//NewDecoder creates a new Decoder instance
func NewDecoder(r io.Reader, header Header) Decoder {
	return &slDecoder{r: r, header: &header}
}

func (d *slDecoder) Next(f FrameReader) error {
	return f.Read(d.r, d.header)
}
