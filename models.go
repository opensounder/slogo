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

func distance(a, b Point) float64 {
	lat1, lon1 := a.GeoLatLon()
	lat2, lon2 := b.GeoLatLon()
	return geoDistance(lat1, lon1, lat2, lon2)
}

func minPoint(a, b Point) Point {
	return Point{
		LatEncoded: min(a.LatEncoded, b.LatEncoded),
		LonEncoded: min(a.LonEncoded, b.LonEncoded),
	}
}

func maxPoint(a, b Point) Point {
	return Point{
		LatEncoded: max(a.LatEncoded, b.LatEncoded),
		LonEncoded: max(a.LonEncoded, b.LonEncoded),
	}
}

func min(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}

func max(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
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
