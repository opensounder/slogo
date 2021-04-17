package slogo

import "math"

const (
	earthRadius         = 6356752.3142
	radConversion       = 180 / math.Pi
	feetMeterConversion = 0.3048
	knotsKphConversion  = 1.85200
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
func Longitude(lon int32) float64 {
	return float64(lon) / earthRadius * radConversion
}

// Convert Lowrance encoded Latitude to decimal degrees. Somewhat expencive procedure.
func Latitude(lat int32) float64 {
	temp := float64(lat) / earthRadius
	temp = math.Exp(temp)
	temp = (2 * math.Atan(temp)) - (math.Pi / 2)
	return temp * radConversion
}
