package goversine

import (
	"math"
	"testing"
)

// Test constants similar to the ones in the original TypeScript tests
const (
	decimalLatitude  = 55.7488
	decimalLongitude = -12.5191

	minLatitude   = -90
	maxLatitude   = 90
	minLongitude  = -180
	maxLongitude  = 180
	minBearing    = 0
	maxBearing    = 359.9999999999
	offset        = 0.0000000001
	earthVMRadius = 6371 // Earth volumetric mean radius in km
)

func TestHaversineDistance(t *testing.T) {
	pointA := MustNewDDPoint(decimalLatitude, decimalLongitude)
	pointB := MustNewDDPoint(decimalLongitude, decimalLatitude)

	// Expected distances in different units
	earthERPointABDistances := map[UnitOfDistance]float64{
		Metre:     9863963.349498631,
		Kilometre: 9863.96334949863,
		Mile:      6129.183297734562,
	}

	// Test different units
	for uod, expected := range earthERPointABDistances {
		haversine := NewHaversine(uod)
		distance := haversine.GetDistance(pointA, pointB)

		// Allow small rounding differences
		if math.Abs(distance-expected) > expected*0.01 {
			t.Errorf("GetDistance with UnitOfDistance %v: got %v, expected %v", uod, distance, expected)
		}
	}

	// Test with custom radius (Earth volumetric mean radius)
	haversine := NewHaversine(Kilometre, earthVMRadius)
	distance := haversine.GetDistance(pointA, pointB)

	expected := 9852.925783760333 // Calculated from TS tests
	if math.Abs(distance-expected) > expected*0.01 {
		t.Errorf("GetDistance with custom radius: got %v, expected %v", distance, expected)
	}
}

func TestBearingCalculation(t *testing.T) {
	pointA := MustNewDDPoint(decimalLatitude, decimalLongitude)
	pointB := MustNewDDPoint(decimalLongitude, decimalLatitude)

	startBearingAB := 114.89155195269666
	endBearingAB := 148.4680591475153

	haversine := NewHaversine(Kilometre)
	bearing := haversine.GetBearing(pointA, pointB)

	if math.Abs(bearing.Start-startBearingAB) > 0.1 {
		t.Errorf("Start bearing: got %v, expected %v", bearing.Start, startBearingAB)
	}

	if math.Abs(bearing.End-endBearingAB) > 0.1 {
		t.Errorf("End bearing: got %v, expected %v", bearing.End, endBearingAB)
	}
}

func TestEndpointCalculation(t *testing.T) {
	pointA := MustNewDDPoint(decimalLatitude, decimalLongitude)
	startBearingAB := 114.89155195269666
	distance := 9863.96334949863 // km

	haversine := NewHaversine(Kilometre)
	endPoint, err := haversine.GetPoint(pointA, startBearingAB, distance)
	if err != nil {
		t.Fatalf("GetPoint returned error: %v", err)
	}

	expectedPointB := MustNewDDPoint(decimalLongitude, decimalLatitude)

	// Allow some rounding differences
	if math.Abs(endPoint.Latitude-expectedPointB.Latitude) > 0.1 ||
		math.Abs(endPoint.Longitude-expectedPointB.Longitude) > 0.1 {
		t.Errorf("GetPoint: got %v, expected %v", endPoint, expectedPointB)
	}
}

func TestPointsSorting(t *testing.T) {
	referencePoint := MustNewDDPoint(0, 0)
	point1 := MustNewDDPoint(1, 1)     // ~157km
	point2 := MustNewDDPoint(0.5, 0.5) // ~79km
	point3 := MustNewDDPoint(2, 2)     // ~314km

	points := []*DDPoint{point1, point3, point2}

	haversine := NewHaversine(Kilometre)

	// Test ascending sort (default)
	sortedPoints := haversine.SortByDistance(referencePoint, points)

	if len(sortedPoints) != 3 {
		t.Fatalf("SortByDistance returned %d points, expected 3", len(sortedPoints))
	}

	// Points should be sorted from nearest to farthest: point2, point1, point3
	if sortedPoints[0] != point2 || sortedPoints[1] != point1 || sortedPoints[2] != point3 {
		t.Errorf("SortByDistance ascending: incorrect order")
	}

	// Test descending sort
	sortedPointsDesc := haversine.SortByDistance(referencePoint, points, Descending)

	// Points should be sorted from farthest to nearest: point3, point1, point2
	if sortedPointsDesc[0] != point3 || sortedPointsDesc[1] != point1 || sortedPointsDesc[2] != point2 {
		t.Errorf("SortByDistance descending: incorrect order")
	}

	// Test empty array
	emptyResult := haversine.SortByDistance(referencePoint, []*DDPoint{})
	if len(emptyResult) != 0 {
		t.Errorf("SortByDistance with empty array: should return empty array")
	}
}

func TestPointsInRange(t *testing.T) {
	referencePoint := MustNewDDPoint(0, 0)
	point1 := MustNewDDPoint(1, 1)     // ~157km
	point2 := MustNewDDPoint(0.5, 0.5) // ~79km
	point3 := MustNewDDPoint(2, 2)     // ~314km

	points := []*DDPoint{point1, point2, point3}

	haversine := NewHaversine(Kilometre)

	// Find points within 100km
	pointsInRange := haversine.GetInRange(referencePoint, points, 100)

	if len(pointsInRange) != 1 || pointsInRange[0] != point2 {
		t.Errorf("GetInRange(100km): expected only point2 to be in range")
	}

	// Find points within 200km
	pointsInRange2 := haversine.GetInRange(referencePoint, points, 200)

	if len(pointsInRange2) != 2 {
		t.Fatalf("GetInRange(200km): expected 2 points to be in range, got %d", len(pointsInRange2))
	}

	// Check if both point1 and point2 are in the results
	found1 := false
	found2 := false
	for _, p := range pointsInRange2 {
		if p == point1 {
			found1 = true
		}
		if p == point2 {
			found2 = true
		}
	}

	if !found1 || !found2 {
		t.Errorf("GetInRange(200km): should include both point1 and point2")
	}

	// Find points within 500km (all points)
	pointsInRange3 := haversine.GetInRange(referencePoint, points, 500)

	if len(pointsInRange3) != 3 {
		t.Errorf("GetInRange(500km): expected all 3 points to be in range")
	}

	// Test negative distance (should return empty array)
	pointsInRangeNeg := haversine.GetInRange(referencePoint, points, -10)

	if len(pointsInRangeNeg) != 0 {
		t.Errorf("GetInRange with negative distance: should return empty array")
	}

	// Test empty array
	pointsInRangeEmpty := haversine.GetInRange(referencePoint, []*DDPoint{}, 100)

	if len(pointsInRangeEmpty) != 0 {
		t.Errorf("GetInRange with empty array: should return empty array")
	}
}

func TestGetInRangeEdgeCases(t *testing.T) {
	haversine := NewHaversine(Kilometre)
	referencePoint := MustNewDDPoint(0, 0)

	t.Run("Nil points slice", func(t *testing.T) {
		var nilSlice []*DDPoint
		result := haversine.GetInRange(referencePoint, nilSlice, 100)
		if len(result) != 0 {
			t.Errorf("Expected empty slice for nil input, got %v items", len(result))
		}
	})

	t.Run("Zero distance", func(t *testing.T) {
		points := []*DDPoint{
			MustNewDDPoint(0.1, 0.1), // Not at reference point
		}
		result := haversine.GetInRange(referencePoint, points, 0)
		if len(result) != 0 {
			t.Errorf("Expected empty slice for zero distance, got %v items", len(result))
		}
	})

	t.Run("Point at exact distance", func(t *testing.T) {
		// Create a point that should be almost exactly 111.2 km from origin (1 degree latitude)
		pointAtExactDistance := MustNewDDPoint(1, 0)
		distance := haversine.GetDistance(referencePoint, pointAtExactDistance)

		points := []*DDPoint{pointAtExactDistance}
		result := haversine.GetInRange(referencePoint, points, distance)
		if len(result) != 1 {
			t.Errorf("Expected 1 point at exact distance boundary, got %v", len(result))
		}

		// Test with slightly smaller distance (should exclude point)
		result = haversine.GetInRange(referencePoint, points, distance-0.1)
		if len(result) != 0 {
			t.Errorf("Expected 0 points just inside boundary, got %v", len(result))
		}
	})
}

func TestSortByDistanceEdgeCases(t *testing.T) {
	haversine := NewHaversine(Kilometre)
	referencePoint := MustNewDDPoint(0, 0)

	t.Run("Nil points slice", func(t *testing.T) {
		var nilSlice []*DDPoint
		result := haversine.SortByDistance(referencePoint, nilSlice)
		if len(result) != 0 {
			t.Errorf("Expected empty slice for nil input, got %v items", len(result))
		}
	})

	t.Run("Single point", func(t *testing.T) {
		singlePoint := MustNewDDPoint(1, 1)
		points := []*DDPoint{singlePoint}
		result := haversine.SortByDistance(referencePoint, points)
		if len(result) != 1 || result[0] != singlePoint {
			t.Errorf("Expected sorted single point to be unchanged")
		}
	})

	t.Run("Points at same distance", func(t *testing.T) {
		// Create points that are equidistant from reference (using latitude only to ensure same distance)
		point1 := MustNewDDPoint(1, 0)
		point2 := MustNewDDPoint(-1, 0)

		points := []*DDPoint{point1, point2}
		result := haversine.SortByDistance(referencePoint, points)
		if len(result) != 2 {
			t.Errorf("Expected 2 points in result, got %v", len(result))
		}

		dist1 := haversine.GetDistance(referencePoint, result[0])
		dist2 := haversine.GetDistance(referencePoint, result[1])
		if math.Abs(dist1-dist2) > 0.0001 {
			t.Errorf("Expected equidistant points to remain equidistant after sorting: %v vs %v", dist1, dist2)
		}
	})
}
