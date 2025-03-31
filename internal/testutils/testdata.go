package testutils

// Sample coordinate data for testing
const (
	// Sample decimal coordinates
	DecimalLatitude  = 55.7488
	DecimalLongitude = -12.5191

	// Sample DMS coordinates for latitude
	DmsLatDegrees = 55.0
	DmsLatMinutes = 44.0
	DmsLatSeconds = 55.68

	// Sample DMS coordinates for longitude
	DmsLongDegrees = -12.0
	DmsLongMinutes = 31.0
	DmsLongSeconds = 8.76

	// Expected bearings between sample points
	StartBearingAB = 114.89155195269666
	EndBearingAB   = 148.4680591475153

	// Expected distances in different units
	DistanceInKm = 9863.96334949863
	DistanceInM  = 9863963.349498631
	DistanceInMi = 6129.183297734562

	// Distance using Earth volumetric mean radius in km
	EarthVMRDistanceInKm = 9852.925783760333
)
