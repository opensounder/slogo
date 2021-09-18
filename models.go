package slogo

import (
	"fmt"
	"io"
)

// Speed in knots
type Speed float32

// Depth in feet
type Depth float32

// Radians angle
type Radians float32

type Flags uint16

const (
	F0 Flags = 1 << iota
	F1
	F2
	F3
	F4
)

// ToKph converts speed to kilometers per hour
func (s Speed) ToKph() float32 {
	return float32(s) * 1.85200
}

// ToMps converts speed to meters per second
func (s Speed) ToMps() float32 {
	return float32(s) * 0.514444
}

// ToMeters convert depth to meters
func (d Depth) ToMeters() float32 {
	return FeetToMeter(float32(d))
}

func (r Radians) ToDeg() float32 {
	return deg32(float32(r))
}

type Point struct {
	YMerc int32
	XMerc int32
}

func (p Point) GeoLatLon() (lat float64, lng float64) {
	return Latitude(p.YMerc), Longitude(p.XMerc)
}

func (p Point) ToGMapsURL(zoom byte) string {
	la, lo := p.GeoLatLon()
	return fmt.Sprintf("https://maps.google.com/maps?q=@%f,%f&z=%d", la, lo, zoom)
}

func (p Point) String() string {
	return fmt.Sprintf("<%d, %d>", p.YMerc, p.XMerc)
}

//Header represents the SLx file header. Same for all formats
type Header struct {
	Format    uint16
	Version   uint16
	Blocksize uint16
	Debug     uint16
}

type Frame interface {
	FrameReader
	Location() Point
}

type FrameReader interface {
	Read(r io.ReadSeeker, header *Header) error
}
