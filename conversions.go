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

// deg32 converts from radians to degrees float32
func deg32(rad float32) float32 {
	return rad * radConversion
}

// deg64 converts from radians to degrees float64
func deg64(rad_val float64) float64 {
	return rad_val * radConversion
}

// FeetToMeter converts
func FeetToMeter(data float32) float32 {
	return data * feetMeterConversion
}

// // KnotsToKph converts
// func KnotsToKph(data float32) float32 {
// 	return data * knotsKphConversion
// }

// Convert Lowrance encoded Longitude to decimal degrees
func Longitude(x int32) float64 {
	return deg64(float64(x) / r_major)
}

// Convert Lowrance encoded Latitude to decimal degrees. Somewhat expencive procedure.
func Latitude(y int32) float64 {
	temp := float64(y) / r_major
	temp = math.Exp(temp)
	temp = (2 * math.Atan(temp)) - halfPi
	return deg64(temp)
}

// rad64 converts degrees to radians float64
func rad64(deg_val float64) float64 {
	return deg_val / radConversion
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
	phi := rad64(lat)
	sinphi := math.Sin(phi)
	con := eccent * sinphi
	com := eccent / 2
	con = math.Pow((1.0-con)/(1.0+con), com)
	ts := math.Tan((math.Pi/2-phi)/2) / con
	y := 0 - r_major*math.Log(ts)
	return int32(y)
}

func merc_x(lon float64) int32 {
	return int32(r_major * rad64(lon))
}

// func merc(lon, lat float64) (x, y int32) {
// 	return merc_x(lon), merc_y(lat)
// }

func has(flags, f Flags) bool { return flags&f != 0 }
