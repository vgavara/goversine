package goversine

import (
	"errors"
	"math"
	"sort"
)

// UnitOfDistance represents the unit of distance for calculations
type UnitOfDistance int

const (
	// Metre represents distance in metres
	Metre UnitOfDistance = iota
	// Kilometre represents distance in kilometres
	Kilometre
	// Mile represents distance in miles
	Mile
)

// Sorting direction for distance-based sorting
type Sorting int

const (
	// Ascending sorts points from nearest to farthest
	Ascending Sorting = iota
	// Descending sorts points from farthest to nearest
	Descending
)

// Earth's equatorial radii in different units
var equatorialEarthRadii = map[UnitOfDistance]float64{
	Metre:     6378137,
	Kilometre: 6378.137,
	Mile:      3963.191,
}

// Haversine is a struct that provides methods for sphere calculations
type Haversine struct {
	// SphereRadius is the radius of the sphere for calculations
	SphereRadius float64
}

// NewHaversine creates a new Haversine calculator with specified unit of distance and optional sphere radius
func NewHaversine(uod UnitOfDistance, sphereRadius ...float64) *Haversine {
	if len(sphereRadius) > 0 && sphereRadius[0] > 0 {
		return &Haversine{SphereRadius: sphereRadius[0]}
	}
	return &Haversine{SphereRadius: equatorialEarthRadii[uod]}
}

// GetDistance calculates the distance between two points using the Haversine formula
func (h *Haversine) GetDistance(pointA, pointB *DDPoint) float64 {
	radALatitude := toRadians(pointA.Latitude)
	radALongitude := toRadians(pointA.Longitude)
	radBLatitude := toRadians(pointB.Latitude)
	radBLongitude := toRadians(pointB.Longitude)

	latitudeDelta := radBLatitude - radALatitude
	longitudeDelta := radBLongitude - radALongitude

	a := math.Pow(math.Sin(latitudeDelta/2), 2) +
		math.Cos(radALatitude)*
			math.Cos(radBLatitude)*
			math.Pow(math.Sin(longitudeDelta/2), 2)
	c := 2 * math.Asin(math.Sqrt(a))

	return c * h.SphereRadius
}

// GetStartBearing calculates the initial bearing from pointA to pointB
func (h *Haversine) getStartBearing(pointA, pointB *DDPoint) float64 {
	radALatitude := toRadians(pointA.Latitude)
	radBLatitude := toRadians(pointB.Latitude)
	deltaLongitude := toRadians(pointB.Longitude - pointA.Longitude)

	y := math.Sin(deltaLongitude) * math.Cos(radBLatitude)
	x := math.Cos(radALatitude)*math.Sin(radBLatitude) -
		math.Sin(radALatitude)*
			math.Cos(radBLatitude)*
			math.Cos(deltaLongitude)
	theta := math.Atan2(y, x)

	return math.Mod(toDegrees(theta)+360, 360)
}

// GetBearing calculates the sphere bearing between two points
func (h *Haversine) GetBearing(startPoint, endPoint *DDPoint) *SphereBearing {
	startBearing := h.getStartBearing(startPoint, endPoint)
	endBearing := math.Mod(h.getStartBearing(endPoint, startPoint)+180, 360)

	return NewSphereBearing(startBearing, endBearing)
}

// GetPoint calculates the end point given a start point, bearing and distance
func (h *Haversine) GetPoint(startPoint *DDPoint, bearing, distance float64) (*DDPoint, error) {
	if distance < 0 {
		return nil, errors.New("distance out of range: must be equal or higher than 0")
	}
	if bearing < 0 || bearing >= 360 {
		return nil, errors.New("bearing out of range: must be between 0 and < 360")
	}

	startLatitude := toRadians(startPoint.Latitude)
	startLongitude := toRadians(startPoint.Longitude)
	angularDistance := distance / h.SphereRadius
	startBearing := toRadians(bearing)

	endLatitude := math.Asin(
		math.Sin(startLatitude)*math.Cos(angularDistance) +
			math.Cos(startLatitude)*
				math.Sin(angularDistance)*
				math.Cos(startBearing),
	)
	endLongitude :=
		startLongitude +
			math.Atan2(
				math.Sin(startBearing)*
					math.Sin(angularDistance)*
					math.Cos(startLatitude),
				math.Cos(angularDistance)-
					math.Sin(startLatitude)*math.Sin(endLatitude),
			)

	return NewDDPoint(toDegrees(endLatitude), toDegrees(endLongitude))
}

// SortByDistance sorts an array of points by their distance to a reference point
func (h *Haversine) SortByDistance(referencePoint *DDPoint, points []*DDPoint, sorting ...Sorting) []*DDPoint {
	if len(points) == 0 {
		return []*DDPoint{}
	}

	sortDirection := Ascending
	if len(sorting) > 0 {
		sortDirection = sorting[0]
	}

	sortedPoints := make([]*DDPoint, len(points))
	copy(sortedPoints, points)

	sort.Slice(sortedPoints, func(i, j int) bool {
		distanceA := h.GetDistance(referencePoint, sortedPoints[i])
		distanceB := h.GetDistance(referencePoint, sortedPoints[j])

		if sortDirection == Ascending {
			return distanceA < distanceB
		}
		return distanceA > distanceB
	})

	return sortedPoints
}

// GetInRange returns points that are within a specified distance from a reference point
func (h *Haversine) GetInRange(referencePoint *DDPoint, points []*DDPoint, distance float64) []*DDPoint {
	if len(points) == 0 || distance < 0 {
		return []*DDPoint{}
	}

	var inRange []*DDPoint
	for _, point := range points {
		if h.GetDistance(referencePoint, point) <= distance {
			inRange = append(inRange, point)
		}
	}

	return inRange
}
