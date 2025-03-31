package goversine

import (
	"math"
	"testing"
)

func TestToRadians(t *testing.T) {
	tests := []struct {
		name    string
		degrees float64
		want    float64
	}{
		{"Zero degrees", 0, 0},
		{"Positive value", 90, math.Pi / 2},
		{"Negative value", -180, -math.Pi},
		{"Arbitrary value", 45, math.Pi / 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := toRadians(tt.degrees)
			if math.Abs(got-tt.want) > 1e-10 {
				t.Errorf("toRadians(%v) = %v, want %v", tt.degrees, got, tt.want)
			}
		})
	}
}

func TestToDegrees(t *testing.T) {
	tests := []struct {
		name    string
		radians float64
		want    float64
	}{
		{"Zero radians", 0, 0},
		{"Positive value", math.Pi / 2, 90},
		{"Negative value", -math.Pi, -180},
		{"Arbitrary value", math.Pi / 4, 45},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := toDegrees(tt.radians)
			if math.Abs(got-tt.want) > 1e-10 {
				t.Errorf("toDegrees(%v) = %v, want %v", tt.radians, got, tt.want)
			}
		})
	}
}

func TestRound(t *testing.T) {
	tests := []struct {
		name          string
		num           float64
		decimalPlaces int
		want          float64
	}{
		{"Round to zero places", 3.14159, 0, 3},
		{"Round to one place", 3.14159, 1, 3.1},
		{"Round to two places", 3.14159, 2, 3.14},
		{"Round up", 3.99999, 0, 4},
		{"Round down", 3.49999, 0, 3},
		{"Round negative number", -3.14159, 2, -3.14},
		{"Round to negative places", 3521.14159, -2, 3500},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := round(tt.num, tt.decimalPlaces)
			if got != tt.want {
				t.Errorf("round(%v, %v) = %v, want %v", tt.num, tt.decimalPlaces, got, tt.want)
			}
		})
	}
}
