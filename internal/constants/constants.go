package constants

// Common constants used across the package
const (
	// Earth constants (in kilometers)
	EarthVolumetricMeanRadius = 6371.0
	EarthEquatorialRadius     = 6378.137

	// Value boundaries
	MinLatitude  = -90.0
	MaxLatitude  = 90.0
	MinLongitude = -180.0
	MaxLongitude = 180.0
	MinBearing   = 0.0
	MaxBearing   = 359.9999999999

	// DMS value boundaries
	MinDegrees = -180.0
	MaxDegrees = 180.0
	MinMinutes = 0.0
	MaxMinutes = 59.0
	MinSeconds = 0.0
	MaxSeconds = 59.9999999999

	// Small offset for boundary testing
	Offset = 0.0000000001

	// Mathematical constants
	RadiansCoefficient = 0.017453292519943295 // Pi/180
)
