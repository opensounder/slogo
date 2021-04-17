package slogo

import "fmt"

type Speed float32
type Depth float32
type Radians float32

func (s Speed) ToKph() float32 {
	return float32(s) * 1.85200
}

func (d Depth) ToMeters() float32 {
	return float32(d) * 0.3048
}

func (r Radians) ToDeg() float32 {
	return RadToDeg(float32(r))
}

//Header represents the log file header
type Header struct {
	Format    uint16
	Version   uint16
	Blocksize uint16
	Reserved1 uint16
}

type Frame interface {
	Location() Point
	GpsSpeed() Speed
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

func (f *FrameV2) Location() Point {
	return Point{f.LatEncoded, f.LonEncoded}
}

type Point struct {
	LatEncoded int32
	LonEncoded int32
}

func (p Point) GeoLatLon() (float64, float64) {
	return Latitude(p.LatEncoded), Longitude(p.LonEncoded)
}

func (p Point) ToGMapsURL(zoom byte) string {
	la, lo := p.GeoLatLon()
	return fmt.Sprintf("https://maps.google.com/maps?q=@%f,%f&z=%d", la, lo, zoom)
}

func (p Point) String() string {
	return fmt.Sprintf("<%d, %d>", p.LatEncoded, p.LonEncoded)
}
