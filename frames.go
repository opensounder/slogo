package slogo

import (
	"encoding/binary"
	"fmt"
	"io"
)

type FrameF2 struct {
	Offset        uint32
	Primary       uint32
	Secondary     uint32
	Down          uint32
	LeftSide      uint32
	RightSide     uint32
	Composite     uint32
	Blocksize     uint16
	LastBlocksize uint16
	Channel       uint16
	Packetsize    uint16
	Frameindex    uint32
	UpperLimit    float32
	LowerLimit    float32
	Reserved1     uint16
	Frequency     uint8
	_             [13]uint8
	WaterDepth    Depth
	KeelDepth     Depth
	_             [28]uint8
	GpsSpeed      Speed
	Temperature   float32
	LonEncoded    int32
	LatEncoded    int32
	WaterSpeed    Speed
	COG           Radians
	Altitude      float32
	Heading       Radians
	Flags         uint16
	_             [6]uint8
	Time          uint32
}

func (f *FrameF2) Location() Point {
	return Point{f.LatEncoded, f.LonEncoded}
}

func (f *FrameF2) Read(r io.Reader, header *Header) error {
	if header.Format != 2 {
		return fmt.Errorf("format %v files is not supported", header.Format)
	}
	err := binary.Read(r, binary.LittleEndian, f)
	if err != nil {
		return err // fmt.Errorf("error reading frame header: %w", err)
	}
	// log.Printf("Offset in hex %x", frame.Offset)
	//TODO Read packet.
	ping := make([]byte, int(f.Packetsize))
	_, err = r.Read(ping)
	if err != nil {
		return fmt.Errorf("error reading frame ping: %w", err)
	}
	// log.Printf("Packetsize: %d, Read: %d\n", frame.Packetsize, n)
	return err
}
