package goversine

import (
	"testing"

	"github.com/vgavara/goversine/internal/constants"
)

func TestSphereBearingValidation(t *testing.T) {
	tests := []struct {
		name      string
		start     float64
		end       float64
		expectErr bool
	}{
		{"Valid minimum values", constants.MinBearing, constants.MinBearing, false},
		{"Valid maximum values", constants.MaxBearing, constants.MaxBearing, false},
		{"Start bearing too small", constants.MinBearing - constants.Offset, constants.MinBearing, true},
		{"Start bearing too large", constants.MaxBearing + constants.Offset, constants.MinBearing, true},
		{"End bearing too small", constants.MinBearing, constants.MinBearing - constants.Offset, true},
		{"End bearing too large", constants.MinBearing, constants.MaxBearing + constants.Offset, true},
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
