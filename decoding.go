package slogo

import (
	"encoding/binary"
	"io"
	"math"
	"os"
)

const (
	EarthRadius    = 6356752.3142
	RadConversion  = 180 / math.Pi
	FeetConversion = 0.3048
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

func Longitude(lon uint32) float64 {
	return float64(lon) / EarthRadius * RadConversion
}

func Latitude(lat uint32) float64 {
	temp := float64(lat) / EarthRadius
	temp = math.Exp(temp)
	temp = (2 * math.Atan(temp)) - (math.Pi / 2)
	return temp * RadConversion
}

func RadToDeg(data float32) float32 {
	return data * RadConversion
}

func FeetToMeter(data float32) float32 {
	return data * FeetConversion
}

func KnotsToKph(data float32) float32 {
	return data * 1.85200
}
