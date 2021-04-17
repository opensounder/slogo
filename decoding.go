package slogo

import (
	"encoding/binary"
	"io"
	"os"
)

//Decoder is the general interface
type Decoder interface {
	DecodeV2(frame *FrameV2) error
}

type slDecoder struct {
	r         io.Reader
	version   uint16
	blocksize uint16
}

//OpenLog is a function
func OpenLog(path string) (*os.File, *Header, error) {
	logfile, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}

	header, err := ReadHeader(logfile)
	if err != nil {
		return nil, nil, err
	}

	return logfile, header, nil
}

//ReadHeader is another function
func ReadHeader(reader io.Reader) (*Header, error) {
	header := Header{}
	if err := binary.Read(reader, binary.LittleEndian, &header.Format); err != nil {
		return nil, err
	}

	if err := binary.Read(reader, binary.LittleEndian, &header.Version); err != nil {
		return nil, err
	}

	if err := binary.Read(reader, binary.LittleEndian, &header.Blocksize); err != nil {
		return nil, err
	}

	if err := binary.Read(reader, binary.LittleEndian, &header.Reserved1); err != nil {
		return nil, err
	}
	return &header, nil
}

//NewDecoder creates a new Decoder instance
func NewDecoder(r io.Reader, version, blocksize uint16) Decoder {
	return &slDecoder{r: r, version: version, blocksize: blocksize}
}

//Decode implements the Decode method
func (d *slDecoder) DecodeV2(frame *FrameV2) error {

	err := binary.Read(d.r, binary.LittleEndian, frame)
	if err != nil {
		return err
	}
	// log.Printf("Offset in hex %x", frame.Offset)
	//TODO Read packet.
	ping := make([]byte, int(frame.Packetsize))
	_, err = d.r.Read(ping)
	if err != nil {
		return err
	}
	// log.Printf("Packetsize: %d, Read: %d\n", frame.Packetsize, n)
	return err
}
