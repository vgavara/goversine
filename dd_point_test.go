package goversine

import (
	"testing"

	"github.com/vgavara/goversine/internal/constants"
	"github.com/vgavara/goversine/internal/testutils"
)

func TestDDPointCreation(t *testing.T) {
	tests := []struct {
		name      string
		latitude  float64
		longitude float64
		expectErr bool
	}{
		{"Valid minimum values", constants.MinLatitude, constants.MinLongitude, false},
		{"Valid maximum values", constants.MaxLatitude, constants.MaxLongitude, false},
		{"Latitude too small", constants.MinLatitude - constants.Offset, 0, true},
		{"Latitude too large", constants.MaxLatitude + constants.Offset, 0, true},
		{"Longitude too small", 0, constants.MinLongitude - constants.Offset, true},
		{"Longitude too large", 0, constants.MaxLongitude + constants.Offset, true},
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
	ddPoint := MustNewDDPoint(testutils.DecimalLatitude, testutils.DecimalLongitude)
	dmsPoint := ddPoint.ToDMSPoint()

	// Test latitude
	if dmsPoint.Latitude.Degrees != testutils.DmsLatDegrees {
		t.Errorf("Latitude degrees: got %v, expected %v", dmsPoint.Latitude.Degrees, testutils.DmsLatDegrees)
	}
	if dmsPoint.Latitude.Minutes != testutils.DmsLatMinutes {
		t.Errorf("Latitude minutes: got %v, expected %v", dmsPoint.Latitude.Minutes, testutils.DmsLatMinutes)
	}
	if round(dmsPoint.Latitude.Seconds, 2) != testutils.DmsLatSeconds {
		t.Errorf("Latitude seconds: got %v, expected %v", round(dmsPoint.Latitude.Seconds, 2), testutils.DmsLatSeconds)
	}

	// Test longitude
	if dmsPoint.Longitude.Degrees != testutils.DmsLongDegrees {
		t.Errorf("Longitude degrees: got %v, expected %v", dmsPoint.Longitude.Degrees, testutils.DmsLongDegrees)
	}
	if dmsPoint.Longitude.Minutes != testutils.DmsLongMinutes {
		t.Errorf("Longitude minutes: got %v, expected %v", dmsPoint.Longitude.Minutes, testutils.DmsLongMinutes)
	}
	if round(dmsPoint.Longitude.Seconds, 2) != testutils.DmsLongSeconds {
		t.Errorf("Longitude seconds: got %v, expected %v", round(dmsPoint.Longitude.Seconds, 2), testutils.DmsLongSeconds)
	}
}
