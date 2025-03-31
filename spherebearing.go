package goversine

import "errors"

// SphereBearing represents the start and end bearings of a path between two points on a sphere
type SphereBearing struct {
	Start float64
	End   float64
}

// NewSphereBearing creates a new sphere bearing
func NewSphereBearing(start, end float64) *SphereBearing {
	return &SphereBearing{
		Start: start,
		End:   end,
	}
}

// ValidateSphereBearing validates bearing values
func ValidateSphereBearing(start, end float64) error {
	if start < 0 || start >= 360 {
		return errors.New("start bearing out of range: it must be between 0 and <360")
	}
	if end < 0 || end >= 360 {
		return errors.New("end bearing out of range: it must be between 0 and <360")
	}
	return nil
}
