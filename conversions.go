package slogo

// https://wiki.openstreetmap.org/wiki/Mercator

import "math"

const (
	// earthRadius = 6356752.3142
	// r_major             = 6378137.000 //Equatorial Radius, WGS84
	r_major             = 6356752.3142 // navico uses polar radius
	r_minor             = 6356752.3142
	radConversion       = 180 / math.Pi
	feetMeterConversion = 0.3048
	knotsKphConversion  = 1.85200
	halfPi              = math.Pi / 2
)

// RadToDeg converts from radians to decimal degrees
func RadToDeg(data float32) float32 {
	return data * radConversion
}

// FeetToMeter converts
func FeetToMeter(data float32) float32 {
	return data * feetMeterConversion
}

// KnotsToKph converts
func KnotsToKph(data float32) float32 {
	return data * knotsKphConversion
}

// Convert Lowrance encoded Longitude to decimal degrees
func Longitude(x int32) float64 {
	return deg(float64(x) / r_major)
}

// Convert Lowrance encoded Latitude to decimal degrees. Somewhat expencive procedure.
func Latitude(y int32) float64 {
	temp := float64(y) / r_major
	temp = math.Exp(temp)
	temp = (2 * math.Atan(temp)) - halfPi
	return deg(temp)
}

func rad(deg_val float64) float64 {
	return deg_val / radConversion
}

func deg(rad_val float64) float64 {
	return rad_val * radConversion
}

func merc_y(lat float64) int32 {
	if lat > 89.5 {
		lat = 89.5
	}
	if lat < -89.5 {
		lat = -89.5
	}
	temp := r_minor / r_major
	eccent := math.Sqrt(1 - math.Pow(temp, 2))
	phi := rad(lat)
	sinphi := math.Sin(phi)
	con := eccent * sinphi
	com := eccent / 2
	con = math.Pow((1.0-con)/(1.0+con), com)
	ts := math.Tan((math.Pi/2-phi)/2) / con
	y := 0 - r_major*math.Log(ts)
	return int32(y)
}

func merc_x(lon float64) int32 {
	return int32(r_major * rad(lon))
}

// func merc(lon, lat float64) (x, y int32) {
// 	return merc_x(lon), merc_y(lat)
// }
