package goversine

import (
	"math"

	"github.com/vgavara/goversine/internal/constants"
)

// toRadians converts a degrees value to radians
func toRadians(degrees float64) float64 {
	return degrees * constants.RadiansCoefficient
}

// toDegrees converts a radians value to degrees
func toDegrees(radians float64) float64 {
	return radians / constants.RadiansCoefficient
}

// round rounds a number to specified decimal places
func round(num float64, decimalPlaces int) float64 {
	powOfTen := math.Pow(10, float64(decimalPlaces))
	return math.Round(num*powOfTen) / powOfTen
}
