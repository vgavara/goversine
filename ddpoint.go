package goversine

import (
	"errors"
	"math"
)

// DDPoint represents a sphere point defined by latitude and longitude in decimal degrees
type DDPoint struct {
	Latitude  float64
	Longitude float64
}

// NewDDPoint creates a new decimal degrees point
func NewDDPoint(latitude, longitude float64) (*DDPoint, error) {
	if latitude < -90 || latitude > 90 || longitude < -180 || longitude > 180 {
		return nil, errors.New("coordinate(s) out of range: they must be in the range -90 to 90 (latitude) or -180 to 180 (longitude)")
	}

	return &DDPoint{
		Latitude:  latitude,
		Longitude: longitude,
	}, nil
}

// MustNewDDPoint creates a new decimal degrees point and panics if coordinates are invalid
func MustNewDDPoint(latitude, longitude float64) *DDPoint {
	point, err := NewDDPoint(latitude, longitude)
	if err != nil {
		panic(err)
	}
	return point
}

// ToDMSPoint converts a decimal degrees point to a degrees-minutes-seconds point
func (p *DDPoint) ToDMSPoint() *DMSPoint {
	latitudeDMS := toDMSCoordinate(p.Latitude)
	longitudeDMS := toDMSCoordinate(p.Longitude)

	return &DMSPoint{
		Latitude:  latitudeDMS,
		Longitude: longitudeDMS,
	}
}

// toDMSCoordinate converts a decimal degrees coordinate to a DMS coordinate
func toDMSCoordinate(ddCoordinate float64) DMSCoordinate {
	absDDCoordinate := math.Abs(ddCoordinate)
	degrees := math.Floor(absDDCoordinate)
	minutes := math.Floor((absDDCoordinate - degrees) * 60)
	seconds := (absDDCoordinate - degrees - minutes/60) * 3600

	var finalDegrees float64
	if ddCoordinate >= 0 {
		finalDegrees = degrees
	} else {
		finalDegrees = -degrees
	}

	return DMSCoordinate{
		Degrees: finalDegrees,
		Minutes: minutes,
		Seconds: seconds,
	}
}
