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
	return float32(d) * 0.3048
}

func (r Radians) ToDeg() float32 {
	return RadToDeg(float32(r))
}

type Point struct {
	YMerc int32
	XMerc int32
}

func PointLatLng(lat float64, lng float64) Point {
	return Point{
		YMerc: merc_y(lat),
		XMerc: merc_x(lng),
	}
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
	Reserved1 uint16
}

type Frame interface {
	FrameReader
	Location() Point
}

type FrameReader interface {
	Read(r io.Reader, header *Header) error
}
