package slogo

import "fmt"

//Header represents the log file header
type Header struct {
	Format    uint16
	Version   uint16
	Blocksize uint16
	Reserved1 uint16
}

func (h *Header) NewFrame() (interface{}, error) {
	switch h.Format {
	case 2:
		return &FrameV2{}, nil
	}
	return nil, fmt.Errorf("invalid header format")
}

type FrameV2 struct {
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
	WaterDepth    float32
	KeelDepth     float32
	_             [28]uint8
	GpsSpeed      float32
	Temperature   float32
	LonEncoded    uint32
	LatEncoded    uint32
	WaterSpeed    float32
	COG           float32
	Altitude      float32
	Heading       float32
	Flags         uint16
	_             [6]uint8
	Time          uint32
}

func (f *FrameV2) Location() Point {
	return Point{f.LatEncoded, f.LonEncoded}
}

type Point struct {
	Lat uint32
	Lon uint32
}

func (p Point) GeoLatLon() (float64, float64) {
	return Latitude(p.Lat), Longitude(p.Lon)
}

type Frame interface {
	Location() Point
}
