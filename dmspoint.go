package goversine

import (
	"errors"
	"math"
)

// DMSCoordinate represents a coordinate in degrees, minutes, seconds format
type DMSCoordinate struct {
	Degrees float64
	Minutes float64
	Seconds float64
}

// NewDMSCoordinate creates a new DMS coordinate
func NewDMSCoordinate(degrees, minutes, seconds float64) (*DMSCoordinate, error) {
	if degrees < -180 || degrees > 180 || minutes < 0 || minutes > 59 || seconds < 0 || seconds >= 60 {
		return nil, errors.New("coordinate(s) out of range: they must be in the range -180 to 180 (degrees) or 0 to 59 (minutes and seconds)")
	}

	return &DMSCoordinate{
		Degrees: degrees,
		Minutes: minutes,
		Seconds: seconds,
	}, nil
}

// MustNewDMSCoordinate creates a new DMS coordinate and panics if values are invalid
func MustNewDMSCoordinate(degrees, minutes, seconds float64) *DMSCoordinate {
	coord, err := NewDMSCoordinate(degrees, minutes, seconds)
	if err != nil {
		panic(err)
	}
	return coord
}

// DMSPoint represents a sphere point defined by latitude and longitude in DMS format
type DMSPoint struct {
	Latitude  DMSCoordinate
	Longitude DMSCoordinate
}

// NewDMSPoint creates a new DMS point
func NewDMSPoint(latitude, longitude DMSCoordinate) (*DMSPoint, error) {
	if latitude.Degrees < -90 || latitude.Degrees > 90 {
		return nil, errors.New("latitude out of range: it must be between -90 and 90")
	}

	return &DMSPoint{
		Latitude:  latitude,
		Longitude: longitude,
	}, nil
}

// MustNewDMSPoint creates a new DMS point and panics if coordinates are invalid
func MustNewDMSPoint(latitude, longitude DMSCoordinate) *DMSPoint {
	point, err := NewDMSPoint(latitude, longitude)
	if err != nil {
		panic(err)
	}
	return point
}

// ToDDPoint converts a DMS point to a decimal degrees point
func (p *DMSPoint) ToDDPoint() *DDPoint {
	ddLatitude := dmsToDD(p.Latitude)
	ddLongitude := dmsToDD(p.Longitude)

	return &DDPoint{
		Latitude:  ddLatitude,
		Longitude: ddLongitude,
	}
}

// dmsToDD converts a DMS coordinate to a decimal degrees coordinate
func dmsToDD(dms DMSCoordinate) float64 {
	absDegrees := math.Abs(dms.Degrees)
	ddValue := absDegrees + dms.Minutes/60 + dms.Seconds/3600

	if dms.Degrees >= 0 {
		return ddValue
	}
	return -ddValue
}
