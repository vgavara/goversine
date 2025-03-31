package goversine

import (
	"testing"
)

func TestDDPointCreation(t *testing.T) {
	tests := []struct {
		name      string
		latitude  float64
		longitude float64
		expectErr bool
	}{
		{"Valid minimum values", minLatitude, minLongitude, false},
		{"Valid maximum values", maxLatitude, maxLongitude, false},
		{"Latitude too small", minLatitude - offset, 0, true},
		{"Latitude too large", maxLatitude + offset, 0, true},
		{"Longitude too small", 0, minLongitude - offset, true},
		{"Longitude too large", 0, maxLongitude + offset, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewDDPoint(tt.latitude, tt.longitude)
			if (err != nil) != tt.expectErr {
				t.Errorf("NewDDPoint(%v, %v) error = %v, expectErr %v",
					tt.latitude, tt.longitude, err, tt.expectErr)
			}
		})
	}
}

func TestDDPointToDMSPoint(t *testing.T) {
	ddPoint := MustNewDDPoint(decimalLatitude, decimalLongitude)
	dmsPoint := ddPoint.ToDMSPoint()

	// Expected DMS values derived from TypeScript tests
	expectedLatDegrees := 55.0
	expectedLatMinutes := 44.0
	expectedLatSeconds := 55.68
	expectedLongDegrees := -12.0
	expectedLongMinutes := 31.0
	expectedLongSeconds := 8.76

	// Test latitude
	if dmsPoint.Latitude.Degrees != expectedLatDegrees {
		t.Errorf("Latitude degrees: got %v, expected %v", dmsPoint.Latitude.Degrees, expectedLatDegrees)
	}
	if dmsPoint.Latitude.Minutes != expectedLatMinutes {
		t.Errorf("Latitude minutes: got %v, expected %v", dmsPoint.Latitude.Minutes, expectedLatMinutes)
	}
	if round(dmsPoint.Latitude.Seconds, 2) != expectedLatSeconds {
		t.Errorf("Latitude seconds: got %v, expected %v", round(dmsPoint.Latitude.Seconds, 2), expectedLatSeconds)
	}

	// Test longitude
	if dmsPoint.Longitude.Degrees != expectedLongDegrees {
		t.Errorf("Longitude degrees: got %v, expected %v", dmsPoint.Longitude.Degrees, expectedLongDegrees)
	}
	if dmsPoint.Longitude.Minutes != expectedLongMinutes {
		t.Errorf("Longitude minutes: got %v, expected %v", dmsPoint.Longitude.Minutes, expectedLongMinutes)
	}
	if round(dmsPoint.Longitude.Seconds, 2) != expectedLongSeconds {
		t.Errorf("Longitude seconds: got %v, expected %v", round(dmsPoint.Longitude.Seconds, 2), expectedLongSeconds)
	}
}
