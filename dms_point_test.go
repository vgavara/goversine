package goversine

import (
	"testing"

	"github.com/vgavara/goversine/internal/constants"
	"github.com/vgavara/goversine/internal/testutils"
)

func TestDMSCoordinateCreation(t *testing.T) {
	tests := []struct {
		name      string
		degrees   float64
		minutes   float64
		seconds   float64
		expectErr bool
	}{
		{"Valid minimum values", constants.MinDegrees, constants.MinMinutes, constants.MinSeconds, false},
		{"Valid maximum values", constants.MaxDegrees, constants.MaxMinutes, constants.MaxSeconds, false},
		{"Degrees too small", constants.MinDegrees - constants.Offset, constants.MinMinutes, constants.MinSeconds, true},
		{"Degrees too large", constants.MaxDegrees + constants.Offset, constants.MinMinutes, constants.MinSeconds, true},
		{"Minutes too small", constants.MinDegrees, constants.MinMinutes - constants.Offset, constants.MinSeconds, true},
		{"Minutes too large", constants.MinDegrees, constants.MaxMinutes + constants.Offset, constants.MinSeconds, true},
		{"Seconds too small", constants.MinDegrees, constants.MinMinutes, constants.MinSeconds - constants.Offset, true},
		{"Seconds too large", constants.MinDegrees, constants.MinMinutes, constants.MaxSeconds + constants.Offset, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewDMSCoordinate(tt.degrees, tt.minutes, tt.seconds)
			if (err != nil) != tt.expectErr {
				t.Errorf("NewDMSCoordinate(%v, %v, %v) error = %v, expectErr %v",
					tt.degrees, tt.minutes, tt.seconds, err, tt.expectErr)
			}
		})
	}
}

func TestDMSPointCreation(t *testing.T) {
	// Create valid coordinates
	validLatitude := DMSCoordinate{Degrees: constants.MinLatitude, Minutes: 0, Seconds: 0}
	validLongitude := DMSCoordinate{Degrees: constants.MinLongitude, Minutes: 0, Seconds: 0}

	// Test valid point creation
	_, err := NewDMSPoint(validLatitude, validLongitude)
	if err != nil {
		t.Errorf("Failed to create valid DMS point: %v", err)
	}

	// Test latitude out of bounds
	invalidLatitude := DMSCoordinate{Degrees: constants.MinLatitude - constants.Offset, Minutes: 0, Seconds: 0}
	_, err = NewDMSPoint(invalidLatitude, validLongitude)
	if err == nil {
		t.Errorf("Should fail when latitude out of bounds")
	}

	invalidLatitude = DMSCoordinate{Degrees: constants.MaxLatitude + constants.Offset, Minutes: 0, Seconds: 0}
	_, err = NewDMSPoint(invalidLatitude, validLongitude)
	if err == nil {
		t.Errorf("Should fail when latitude out of bounds")
	}
}

func TestDMSPointToDDPoint(t *testing.T) {
	// DMS coordinates for testing
	dmsLatitude := DMSCoordinate{
		Degrees: testutils.DmsLatDegrees,
		Minutes: testutils.DmsLatMinutes,
		Seconds: testutils.DmsLatSeconds,
	}

	dmsLongitude := DMSCoordinate{
		Degrees: testutils.DmsLongDegrees,
		Minutes: testutils.DmsLongMinutes,
		Seconds: testutils.DmsLongSeconds,
	}

	dmsPoint := DMSPoint{
		Latitude:  dmsLatitude,
		Longitude: dmsLongitude,
	}

	ddPoint := dmsPoint.ToDDPoint()

	// Check conversion accuracy
	if round(ddPoint.Latitude, 4) != testutils.DecimalLatitude {
		t.Errorf("Latitude conversion: got %v, expected %v", round(ddPoint.Latitude, 4), testutils.DecimalLatitude)
	}

	if round(ddPoint.Longitude, 4) != testutils.DecimalLongitude {
		t.Errorf("Longitude conversion: got %v, expected %v", round(ddPoint.Longitude, 4), testutils.DecimalLongitude)
	}
}
