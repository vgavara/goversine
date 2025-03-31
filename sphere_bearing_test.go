package goversine

import (
	"testing"
)

func TestSphereBearingValidation(t *testing.T) {
	tests := []struct {
		name      string
		start     float64
		end       float64
		expectErr bool
	}{
		{"Valid minimum values", minBearing, minBearing, false},
		{"Valid maximum values", maxBearing, maxBearing, false},
		{"Start bearing too small", minBearing - offset, minBearing, true},
		{"Start bearing too large", maxBearing + offset, minBearing, true},
		{"End bearing too small", minBearing, minBearing - offset, true},
		{"End bearing too large", minBearing, maxBearing + offset, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateSphereBearing(tt.start, tt.end)
			if (err != nil) != tt.expectErr {
				t.Errorf("ValidateSphereBearing(%v, %v) error = %v, expectErr %v",
					tt.start, tt.end, err, tt.expectErr)
			}
		})
	}
}

func TestSphereBearingCreation(t *testing.T) {
	start := 45.0
	end := 225.0

	bearing := NewSphereBearing(start, end)

	if bearing.Start != start {
		t.Errorf("Start bearing: got %v, expected %v", bearing.Start, start)
	}

	if bearing.End != end {
		t.Errorf("End bearing: got %v, expected %v", bearing.End, end)
	}
}
