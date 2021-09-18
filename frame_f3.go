package slogo

import (
	"encoding/binary"
	"fmt"
	"io"
)

type FrameF3Info struct {
	Offset        uint32 // I
	TBD1          uint32
	Framesize     uint16
	LastFramesize uint16
	Channel       uint32
	Frameindex    uint32
	UpperLimit    float32
	LowerLimit    float32
	_             [12]uint8
	TimeStamp     uint32
	Payloadsize   uint32
	WaterDepth    Depth
	Frequency     uint32
	_             [28]uint8
	GpsSpeed      Speed
	Temperature   float32
	XMerc         int32
	YMerc         int32
	WaterSpeed    Speed
	COG           Radians
	Altitude      float32
	Heading       Radians
	Flags         Flags
	_             [6]uint8
	Time1         uint32
}

type FrameF3 struct {
	FrameF3Info
	Payload []byte
}

func (f *FrameF3) Location() Point {
	return Point{f.XMerc, f.YMerc}
}

func (f *FrameF3) Read(r io.ReadSeeker, header *Header) error {
	if header.Format != 3 {
		return fmt.Errorf("format %v files is not supported", header.Format)
	}
	info := FrameF3Info{}
	err := binary.Read(r, binary.LittleEndian, &info)
	if err != nil {
		return err
	}
	f.FrameF3Info = info
	payloadsize := f.Payloadsize
	if header.Version == 1 && !has(f.Flags, F3) {
		payloadsize = uint32(f.Framesize) - 168
	}

	if header.Version == 1 || (header.Version == 2 && f.Channel <= 5) {
		// TODO: what the hell are these extry bytes?
		extra := 168 - int64(binary.Size(info))
		r.Seek(extra, io.SeekCurrent)
	}

	// Read payload.
	payload := make([]byte, int(payloadsize))
	_, err = r.Read(payload)
	if err != nil {
		return fmt.Errorf("error reading frame ping: %w", err)
	}
	f.Payload = payload
	return err
}
