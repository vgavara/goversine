# goversine

![License:](https://img.shields.io/github/license/vgavara/goversine)

## About

_goversine_ is a Golang package for distance calculation on a sphere surface with both decimal degrees (DD) or degrees minutes seconds (DMS) coordinates.

_goversine_ main features:

- Allows calculating distances between to points in metres, kilometres or miles.
- Allows using decimal degree (DD) or degrees minutes seconds (DMS) coordinates.
- Allows calculating start and end bearings of a path between two points.
- Allows calculating the end point of a path given its start point, start bearing and distance.
- Allows sorting points by distance from a reference point.
- Allows finding points within a specific distance range from a reference point.

## At a glance

### Calculate the distance, in kilometres, between two decimal degrees coordinates

```go
package main

import (
	"fmt"
	
	"github.com/vgavara/goversine"
)

func main() {
	// New York decimal degrees (DD) coordinates are:
	// latitude 40.730610 N,
	// longitude 73.935242 W
	// Using error-returning constructor
	newYork, err := goversine.NewDDPoint(40.73061, -73.935242)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Madrid decimal degrees (DD) coordinates are:
	// latitude 40.416775 N,
	// longitude 3.703790 W
	// Using panic-on-error constructor for known-good values
	madrid := goversine.MustNewDDPoint(40.416775, -3.70379)

	haversine := goversine.NewHaversine(goversine.Kilometre)
	distance := haversine.GetDistance(newYork, madrid)

	fmt.Printf("The distance from New York to Madrid is %.2f kilometres.\n", distance)
}
```

### Calculate the distance, in miles, between two degrees minutes seconds (DMS) coordinates

```go
package main

import (
	"fmt"
	
	"github.com/vgavara/goversine"
)

func main() {
	// New York DMS coordinates are
	// latitude 40° 43' 50.1960'' N,
	// longitude 73° 56' 6.8712'' W
	// Using error-checking constructors
	newYorkLat, err := goversine.NewDMSCoordinate(40, 43, 50.196)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	newYorkLong, err := goversine.NewDMSCoordinate(-73, 56, 6.8712)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	newYork, err := goversine.NewDMSPoint(*newYorkLat, *newYorkLong)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Madrid DMS coordinates are
	// latitude 40° 25' 0.3900'' N,
	// longitude 3° 42' 13.6440'' W
	// Using panic-on-error constructors for known-good values
	madridLat := goversine.MustNewDMSCoordinate(40, 25, 0.39)
	madridLong := goversine.MustNewDMSCoordinate(-3, 42, 13.644)
	madrid := goversine.MustNewDMSPoint(*madridLat, *madridLong)

	haversine := goversine.NewHaversine(goversine.Mile)
	distance := haversine.GetDistance(newYork.ToDDPoint(), madrid.ToDDPoint())

	fmt.Printf("The distance from New York to Madrid is %.2f miles.\n", distance)
}
```

### Calculate the bearing between two points

```go
package main

import (
	"fmt"
	
	"github.com/vgavara/goversine"
)

func main() {
	newYork := goversine.MustNewDDPoint(40.73061, -73.935242)
	madrid := goversine.MustNewDDPoint(40.416775, -3.70379)

	haversine := goversine.NewHaversine(goversine.Kilometre)
	bearing := haversine.GetBearing(newYork, madrid)

	fmt.Printf(
		"The start bearing of the path from New York to Madrid is %.2f degrees, and the end bearing is %.2f degrees.\n",
		bearing.Start, bearing.End,
	)
}
```

### Calculate the endpoint of a path

```go
package main

import (
	"fmt"
	
	"github.com/vgavara/goversine"
)

func main() {
	newYork := goversine.MustNewDDPoint(40.73061, -73.935242)
	bearing := 65.71472
	distance := 5762.0

	haversine := goversine.NewHaversine(goversine.Kilometre)
	madrid, err := haversine.GetPoint(newYork, bearing, distance)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf(
		"Madrid is the endpoint of the path starting in New York with a bearing of %.2f degrees at a distance of %.2f kilometers\n",
		bearing, distance,
	)
	fmt.Printf("Madrid coordinates: Latitude %.6f, Longitude %.6f\n", madrid.Latitude, madrid.Longitude)
}
```

### Sort points by distance from a reference point

```go
package main

import (
	"fmt"
	
	"github.com/vgavara/goversine"
)

func main() {
	madrid := goversine.MustNewDDPoint(40.416775, -3.70379)
	
	// Create a slice to hold the cities
	cities := make([]*goversine.DDPoint, 0, 4)
	
	// Add Berlin using error-checking constructor
	berlin, err := goversine.NewDDPoint(52.5200, 13.4050)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		cities = append(cities, berlin)
	}
	
	// Add remaining cities using panic-on-error constructor for brevity
	cities = append(cities, goversine.MustNewDDPoint(40.73061, -73.935242)) // New York
	cities = append(cities, goversine.MustNewDDPoint(48.8566, 2.3522))      // Paris
	cities = append(cities, goversine.MustNewDDPoint(35.6762, 139.6503))    // Tokyo

	haversine := goversine.NewHaversine(goversine.Kilometre)
	sortedCities := haversine.SortByDistance(madrid, cities)

	fmt.Println("The cities sorted by distance from Madrid (nearest to farthest) are:")
	for i, point := range sortedCities {
		fmt.Printf("%d. Latitude %.6f, Longitude %.6f\n", i+1, point.Latitude, point.Longitude)
	}
}
```

### Find points within a specific distance range

```go
package main

import (
	"fmt"
	
	"github.com/vgavara/goversine"
)

func main() {
	madrid := goversine.MustNewDDPoint(40.416775, -3.70379)
	cities := []*goversine.DDPoint{
		goversine.MustNewDDPoint(52.5200, 13.4050),     // Berlin
		goversine.MustNewDDPoint(40.73061, -73.935242), // New York
		goversine.MustNewDDPoint(48.8566, 2.3522),      // Paris
		goversine.MustNewDDPoint(35.6762, 139.6503),    // Tokyo
	}

	haversine := goversine.NewHaversine(goversine.Kilometre)
	// Find cities within 2000 kilometers of Madrid
	citiesInRange := haversine.GetInRange(madrid, cities, 2000)

	fmt.Println("The cities within 2000 kilometers from Madrid are:")
	for i, point := range citiesInRange {
		fmt.Printf("%d. Latitude %.6f, Longitude %.6f\n", i+1, point.Latitude, point.Longitude)
	}
}
```

## Installation

To install the package, use the go get command:

```shell
go get github.com/vgavara/goversine
```

## Usage

### Overview

The [`Haversine`](#Haversine) struct supports the implementation of the sphere path resolvers (distance and bearing between two points, end point given start point, bearing and distance).

It uses as input decimal degrees (DD) coordinates defined as [`DDPoint`](#DDPoint) struct instances, that can be converted into degrees minutes seconds (DMS) coordinates as instances of the [`DMSPoint`](#DMSPoint) struct. Each [`DMSPoint`](#DMSPoint) object instance is composed by two [`DMSCoordinate`](#DMSCoordinate) struct instances.

The [`SphereBearing`](#SphereBearing) struct represents a tuple with the start and end bearings of a sphere path (orthodrome) between two points.

<a name="DDPoint"></a>

### DDPoint struct

Sphere point defined by a latitude and a longitude in decimal degrees (DD)

- Constructor functions:
  - [NewDDPoint(latitude, longitude)](#NewDDPoint) - Returns point and error
  - [MustNewDDPoint(latitude, longitude)](#MustNewDDPoint) - Panics on invalid input
- Fields:
  - [Latitude](#DDPoint+Latitude)
  - [Longitude](#DDPoint+Longitude)
- Methods:
  - [.ToDMSPoint()](#DDPoint+ToDMSPoint) ⇒ [`DMSPoint`](#DMSPoint)

<a name="NewDDPoint"></a>

#### NewDDPoint(latitude, longitude) (*DDPoint, error)

Creates a sphere point object instance with error handling.

```go
point, err := goversine.NewDDPoint(40.73061, -73.935242)
if err != nil {
    // Handle error
    return
}
fmt.Printf("Decimal coordinates: Latitude %.6f, Longitude %.6f\n", point.Latitude, point.Longitude)
```

<a name="MustNewDDPoint"></a>

#### MustNewDDPoint(latitude, longitude) *DDPoint

Creates a sphere point object instance, panics if coordinates are invalid. Use this when you're certain the input is valid or during initialization where a failure should stop the program.

```go
point := goversine.MustNewDDPoint(40.73061, -73.935242)
fmt.Printf("Decimal coordinates: Latitude %.6f, Longitude %.6f\n", point.Latitude, point.Longitude)
```

<a name="DDPoint+Latitude"></a>

#### ddPoint.Latitude ⇒ `float64`

The point latitude set when creating the struct.

<a name="DDPoint+Longitude"></a>

#### ddPoint.Longitude ⇒ `float64`

The point longitude set when creating the struct.

<a name="DDPoint+ToDMSPoint"></a>

#### ddPoint.ToDMSPoint() ⇒ [`DMSPoint`](#DMSPoint)

Gets the equivalent point in degrees minutes seconds (DMS) notation.

```go
point := goversine.MustNewDDPoint(40.73061, -73.935242)
dmsPoint := point.ToDMSPoint()

fmt.Printf(
  "DMS coordinates: Latitude %v° %v' %v\", Longitude %v° %v' %v\"\n",
  dmsPoint.Latitude.Degrees, dmsPoint.Latitude.Minutes, dmsPoint.Latitude.Seconds,
  dmsPoint.Longitude.Degrees, dmsPoint.Longitude.Minutes, dmsPoint.Longitude.Seconds,
)
```

<a name="DMSCoordinate"></a>

### DMSCoordinate struct

Latitude/Longitude coordinate defined in degrees minutes seconds (DMS).

- Constructor functions:
  - [NewDMSCoordinate(degrees, minutes, seconds)](#NewDMSCoordinate) - Returns coordinate and error
  - [MustNewDMSCoordinate(degrees, minutes, seconds)](#MustNewDMSCoordinate) - Panics on invalid input
- Fields:
  - [Degrees](#DMSCoordinate+Degrees)
  - [Minutes](#DMSCoordinate+Minutes)
  - [Seconds](#DMSCoordinate+Seconds)

<a name="NewDMSCoordinate"></a>

#### NewDMSCoordinate(degrees, minutes, seconds) (*DMSCoordinate, error)

Creates a DMS coordinate struct with error handling.

```go
coordinate, err := goversine.NewDMSCoordinate(40, 43, 50.196)
if err != nil {
    // Handle error
    return
}
fmt.Printf("DMS: %v° %v' %v\"\n", coordinate.Degrees, coordinate.Minutes, coordinate.Seconds)
```

<a name="MustNewDMSCoordinate"></a>

#### MustNewDMSCoordinate(degrees, minutes, seconds) *DMSCoordinate

Creates a DMS coordinate struct, panics if values are invalid. Use this when you're certain the input is valid or during initialization where a failure should stop the program.

```go
coordinate := goversine.MustNewDMSCoordinate(40, 43, 50.196)
fmt.Printf("DMS: %v° %v' %v\"\n", coordinate.Degrees, coordinate.Minutes, coordinate.Seconds)
```

<a name="DMSCoordinate+Degrees"></a>

#### dmsCoordinate.Degrees ⇒ `float64`

The degrees component of the coordinate.

<a name="DMSCoordinate+Minutes"></a>

#### dmsCoordinate.Minutes ⇒ `float64`

The minutes component of the coordinate.

<a name="DMSCoordinate+Seconds"></a>

#### dmsCoordinate.Seconds ⇒ `float64`

The seconds component of the coordinate.

<a name="DMSPoint"></a>

### DMSPoint struct

Sphere point defined by a latitude and a longitude in degrees minutes seconds (DMS).

- Constructor functions:
  - [NewDMSPoint(latitude, longitude)](#NewDMSPoint) - Returns point and error
  - [MustNewDMSPoint(latitude, longitude)](#MustNewDMSPoint) - Panics on invalid input
- Fields:
  - [Latitude](#DMSPoint+Latitude)
  - [Longitude](#DMSPoint+Longitude)
- Methods:
  - [.ToDDPoint()](#DMSPoint+ToDDPoint) ⇒ [`DDPoint`](#DDPoint)

<a name="NewDMSPoint"></a>

#### NewDMSPoint(latitude, longitude) (*DMSPoint, error)

Creates a sphere point struct in DMS notation with error handling.

```go
latCoord, err := goversine.NewDMSCoordinate(40, 43, 50.196)
if err != nil {
    // Handle error
    return
}
longCoord, err := goversine.NewDMSCoordinate(-73, 56, 6.8712)
if err != nil {
    // Handle error
    return
}

point, err := goversine.NewDMSPoint(*latCoord, *longCoord)
if err != nil {
    // Handle error
    return
}
```

<a name="MustNewDMSPoint"></a>

#### MustNewDMSPoint(latitude, longitude) *DMSPoint

Creates a sphere point struct in DMS notation, panics if coordinates are invalid. Use this when you're certain the input is valid or during initialization where a failure should stop the program.

```go
latCoord := goversine.MustNewDMSCoordinate(40, 43, 50.196)
longCoord := goversine.MustNewDMSCoordinate(-73, 56, 6.8712)

point := goversine.MustNewDMSPoint(*latCoord, *longCoord)
```

<a name="DMSPoint+Latitude"></a>

#### dmsPoint.Latitude ⇒ `DMSCoordinate`

The latitude coordinate in DMS format.

<a name="DMSPoint+Longitude"></a>

#### dmsPoint.Longitude ⇒ `DMSCoordinate`

The longitude coordinate in DMS format.

<a name="DMSPoint+ToDDPoint"></a>

#### dmsPoint.ToDDPoint() ⇒ `DDPoint`

Gets the equivalent point in decimal degrees notation.

```go
latCoord := goversine.MustNewDMSCoordinate(40, 43, 50.196)
longCoord := goversine.MustNewDMSCoordinate(-73, 56, 6.8712)
dmsPoint := goversine.MustNewDMSPoint(*latCoord, *longCoord)

ddPoint := dmsPoint.ToDDPoint()
fmt.Printf("DD coordinates: Latitude %.6f, Longitude %.6f\n", ddPoint.Latitude, ddPoint.Longitude)
```

<a name="Haversine"></a>

### Haversine struct

Haversine formula resolver.

- Constructor function:
  - [NewHaversine(uod, [sphereRadius])](#NewHaversine)
- Methods:
  - [.GetBearing(startPoint, endPoint)](#Haversine+GetBearing) ⇒ [`SphereBearing`](#SphereBearing)
  - [.GetDistance(pointA, pointB)](#Haversine+GetDistance) ⇒ `float64`
  - [.GetInRange(referencePoint, points, distance)](#Haversine+GetInRange) ⇒ `[]*DDPoint`
  - [.GetPoint(startPoint, bearing, distance)](#Haversine+GetPoint) ⇒ [`DDPoint`](#DDPoint)
  - [.SortByDistance(referencePoint, points, [sorting])](#Haversine+SortByDistance) ⇒ `[]*DDPoint`

<a name="NewHaversine"></a>

#### NewHaversine(uod, [sphereRadius]) *Haversine

Initializes the Haversine resolver.

```go
// Basic usage with default Earth radius
haversine := goversine.NewHaversine(goversine.Mile)

// With custom sphere radius
customRadius := 6371.0 // Earth volumetric mean radius in km
haversine := goversine.NewHaversine(goversine.Kilometre, customRadius)
```

<a name="Haversine+GetBearing"></a>

#### haversine.GetBearing(startPoint, endPoint) *SphereBearing

Calculates the sphere bearing between two points.

```go
newYork := goversine.MustNewDDPoint(40.73061, -73.935242)
madrid := goversine.MustNewDDPoint(40.416775, -3.70379)

haversine := goversine.NewHaversine(goversine.Kilometre)
bearing := haversine.GetBearing(newYork, madrid)

fmt.Printf("Start bearing: %.2f°, End bearing: %.2f°\n", bearing.Start, bearing.End)
```

<a name="Haversine+GetDistance"></a>

#### haversine.GetDistance(pointA, pointB) float64

Calculates the distance between two sphere points.

```go
newYork := goversine.MustNewDDPoint(40.73061, -73.935242)
madrid := goversine.MustNewDDPoint(40.416775, -3.70379)

haversine := goversine.NewHaversine(goversine.Kilometre)
distance := haversine.GetDistance(newYork, madrid)

fmt.Printf("Distance: %.2f km\n", distance)
```

<a name="Haversine+GetInRange"></a>

#### haversine.GetInRange(referencePoint, points, distance) []*DDPoint

Finds points within a specified distance from a reference point.

```go
madrid := goversine.MustNewDDPoint(40.416775, -3.70379)
cities := []*goversine.DDPoint{
    goversine.MustNewDDPoint(52.5200, 13.4050),     // Berlin
    goversine.MustNewDDPoint(40.73061, -73.935242), // New York
    goversine.MustNewDDPoint(48.8566, 2.3522),      // Paris
    goversine.MustNewDDPoint(35.6762, 139.6503),    // Tokyo
}

haversine := goversine.NewHaversine(goversine.Kilometre)
// Find cities within 2000 kilometers of Madrid
nearCities := haversine.GetInRange(madrid, cities, 2000)

fmt.Printf("Found %d cities within 2000km of Madrid\n", len(nearCities))
```

<a name="Haversine+GetPoint"></a>

#### haversine.GetPoint(startPoint, bearing, distance) (*DDPoint, error)

Calculates the end point given a start point, bearing and distance.

```go
newYork := goversine.MustNewDDPoint(40.73061, -73.935242)
bearing := 65.71472
distance := 5762.0

haversine := goversine.NewHaversine(goversine.Kilometre)
endpoint, err := haversine.GetPoint(newYork, bearing, distance)
if err != nil {
    // Handle error
    return
}

fmt.Printf("Endpoint: Latitude %.6f, Longitude %.6f\n", endpoint.Latitude, endpoint.Longitude)
```

<a name="Haversine+SortByDistance"></a>

#### haversine.SortByDistance(referencePoint, points, [sorting]) []*DDPoint

Sorts points by distance from a reference point.

```go
madrid := goversine.MustNewDDPoint(40.416775, -3.70379)
cities := []*goversine.DDPoint{
    goversine.MustNewDDPoint(52.5200, 13.4050),     // Berlin
    goversine.MustNewDDPoint(40.73061, -73.935242), // New York
    goversine.MustNewDDPoint(48.8566, 2.3522),      // Paris
    goversine.MustNewDDPoint(35.6762, 139.6503),    // Tokyo
}

haversine := goversine.NewHaversine(goversine.Kilometre)
// Sort ascending (nearest to farthest)
nearestCities := haversine.SortByDistance(madrid, cities)
// Sort descending (farthest to nearest)
farthestCities := haversine.SortByDistance(madrid, cities, goversine.Descending)
```

<a name="SphereBearing"></a>

### SphereBearing struct

Sphere bearing representing start and end bearings of a path between two points.

- Constructor function:
  - [NewSphereBearing(start, end)](#NewSphereBearing) ⇒ `*SphereBearing`
  - [ValidateSphereBearing(start, end)](#ValidateSphereBearing) ⇒ `error`
- Fields:
  - [Start](#SphereBearing+Start)
  - [End](#SphereBearing+End)

<a name="NewSphereBearing"></a>

#### NewSphereBearing(start, end) *SphereBearing

Creates a new sphere bearing struct.

```go
bearing := goversine.NewSphereBearing(60.5, 181.0)
fmt.Printf("Start bearing: %.2f°, End bearing: %.2f°\n", bearing.Start, bearing.End)
```

<a name="ValidateSphereBearing"></a>

#### ValidateSphereBearing(start, end) error

Validates bearing values.

```go
err := goversine.ValidateSphereBearing(60.5, 181.0)
if err != nil {
    // Handle error
    return
}
```

<a name="SphereBearing+Start"></a>

#### sphereBearing.Start ⇒ `float64`

The start bearing value.

<a name="SphereBearing+End"></a>

#### sphereBearing.End ⇒ `float64`

The end bearing value.

<a name="UnitOfDistance"></a>

### UnitOfDistance type

```go
type UnitOfDistance int

const (
	Metre     UnitOfDistance = iota // Distance in metres
	Kilometre                        // Distance in kilometres
	Mile                             // Distance in miles
)
```

<a name="Sorting"></a>

### Sorting type

```go
type Sorting int

const (
	Ascending  Sorting = iota // Sort points from nearest to farthest
	Descending               // Sort points from farthest to nearest
)
```

## Support

In order to notify some problem or suggest an improvement or new feature, submit an issue in the GitHub repository [issues](https://github.com/vgavara/goversine/issues) section.

## License

This package is licensed under the [MIT](https://opensource.org/licenses/MIT) terms of use.

## Contact

You can contact the package creator via [email](mailto:vgavara@gmail.com), [GitHub](https://github.com/vgavara) or [LinkedIn](https://www.linkedin.com/in/vgavara/).
