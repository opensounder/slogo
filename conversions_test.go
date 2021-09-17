package slogo

import (
	"math"
	"testing"
)

func Test_x_Longitude(t *testing.T) {
	tests := []struct {
		name string
		x    int32
		lon  float64
	}{
		// TODO: Add test cases.
		{"a", 1373465, 12.379552312136807},
		{"b", 0, 0},
		{"c", 1, 9.013372974292616e-06},
		{"max", math.MaxInt32, 19356.071066605145},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Longitude(tt.x); !almostEqual(got, tt.lon, float64Epsilon) {
				t.Errorf("Longitude() = %v, want %v", got, tt.lon)
			}
			if got_x := merc_x(tt.lon); got_x != tt.x {
				t.Errorf("merc_x() = %v, want %v", got_x, tt.x)
			}
		})
	}
}

func Test_y_Latitude(t *testing.T) {

	tests := []struct {
		name string
		y    int32
		lat  float64
	}{
		// TODO: Add test cases.
		{"zero", 0, 0.0},
		{"one", 1, 9.013372970355275e-06},
		{"a", 8180800, 59.12899916049587},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Latitude(tt.y); !almostEqual(got, tt.lat, float64Epsilon) {
				t.Errorf("Latitude() = %v, want %v", got, tt.lat)
			}
			if got_y := merc_y(tt.lat); got_y != tt.y {
				t.Errorf("merc_y() = %v, want %v", got_y, tt.y)
			}
		})
	}
}
