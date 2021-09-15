package slogo

import (
	"fmt"
	"io"
)

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
