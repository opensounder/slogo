package slogo

import (
	"encoding/binary"
	"fmt"
	"io"
)

type FrameF2 struct {
	Offset            uint32
	PreviousPrimary   uint32
	PreviousSecondary uint32
	PreviousDownscan  uint32
	PreviousLeftSide  uint32
	PreviousRightSide uint32
	PreviousComposite uint32
	Framesize         uint16
	LastFramesize     uint16
	Channel           uint16
	Packetsize        uint16
	Frameindex        uint32
	UpperLimit        float32
	LowerLimit        float32
	_                 uint16
	Frequency         uint8
	_                 [13]uint8
	WaterDepth        Depth
	KeelDepth         Depth
	_                 [28]uint8
	GpsSpeed          Speed
	Temperature       float32
	XMerc             int32
	YMerc             int32
	WaterSpeed        Speed
	COG               Radians
	Altitude          float32
	Heading           Radians
	Flags             uint16
	_                 [6]uint8
	Time1             uint32
}

func (f *FrameF2) Location() Point {
	return Point{f.YMerc, f.XMerc}
}

func (f *FrameF2) Read(r io.ReadSeeker, header *Header) error {
	if header.Format != 2 {
		return fmt.Errorf("format %v files is not supported", header.Format)
	}
	err := binary.Read(r, binary.LittleEndian, f)
	if err != nil {
		// fmt.Printf("error %v. %+v", err, f)
		return err // fmt.Errorf("error reading frame header: %w", err)
	}

	if f.Packetsize > header.Blocksize {
		return fmt.Errorf("packetsize %v > %v", f.Packetsize, header.Blocksize)
	}
	ping := make([]byte, int(f.Packetsize))
	_, err = r.Read(ping)
	if err != nil {
		return fmt.Errorf("error reading frame ping: %w", err)
	}
	// log.Printf("Packetsize: %d, Read: %d\n", frame.Packetsize, n)
	return err
}
