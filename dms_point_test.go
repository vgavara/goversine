package goversine

import (
	"testing"
)

func TestDMSCoordinateCreation(t *testing.T) {
	// Minimum and maximum values
	minDegrees := -180.0
	maxDegrees := 180.0
	minMinutes := 0.0
	maxMinutes := 59.0
	minSeconds := 0.0
	maxSeconds := 59.9999999999

	tests := []struct {
		name      string
		degrees   float64
		minutes   float64
		seconds   float64
		expectErr bool
	}{
		{"Valid minimum values", minDegrees, minMinutes, minSeconds, false},
		{"Valid maximum values", maxDegrees, maxMinutes, maxSeconds, false},
		{"Degrees too small", minDegrees - offset, minMinutes, minSeconds, true},
		{"Degrees too large", maxDegrees + offset, minMinutes, minSeconds, true},
		{"Minutes too small", minDegrees, minMinutes - offset, minSeconds, true},
		{"Minutes too large", minDegrees, maxMinutes + offset, minSeconds, true},
		{"Seconds too small", minDegrees, minMinutes, minSeconds - offset, true},
		{"Seconds too large", minDegrees, minMinutes, maxSeconds + offset, true},
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
	validLatitude := DMSCoordinate{Degrees: minLatitude, Minutes: 0, Seconds: 0}
	validLongitude := DMSCoordinate{Degrees: minLongitude, Minutes: 0, Seconds: 0}

	// Test valid point creation
	_, err := NewDMSPoint(validLatitude, validLongitude)
	if err != nil {
		t.Errorf("Failed to create valid DMS point: %v", err)
	}

	// Test latitude out of bounds
	invalidLatitude := DMSCoordinate{Degrees: minLatitude - offset, Minutes: 0, Seconds: 0}
	_, err = NewDMSPoint(invalidLatitude, validLongitude)
	if err == nil {
		t.Errorf("Should fail when latitude out of bounds")
	}

	invalidLatitude = DMSCoordinate{Degrees: maxLatitude + offset, Minutes: 0, Seconds: 0}
	_, err = NewDMSPoint(invalidLatitude, validLongitude)
	if err == nil {
		t.Errorf("Should fail when latitude out of bounds")
	}
}

func TestDMSPointToDDPoint(t *testing.T) {
	// DMS coordinates from original tests
	dmsLatitude := DMSCoordinate{
		Degrees: 55,
		Minutes: 44,
		Seconds: 55.68,
	}

	dmsLongitude := DMSCoordinate{
		Degrees: -12,
		Minutes: 31,
		Seconds: 8.76,
	}

	dmsPoint := DMSPoint{
		Latitude:  dmsLatitude,
		Longitude: dmsLongitude,
	}

	ddPoint := dmsPoint.ToDDPoint()

	// Check conversion accuracy
	if round(ddPoint.Latitude, 4) != decimalLatitude {
		t.Errorf("Latitude conversion: got %v, expected %v", round(ddPoint.Latitude, 4), decimalLatitude)
	}

	if round(ddPoint.Longitude, 4) != decimalLongitude {
		t.Errorf("Longitude conversion: got %v, expected %v", round(ddPoint.Longitude, 4), decimalLongitude)
	}
}
