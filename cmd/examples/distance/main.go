package main

import (
	"fmt"

	"github.com/vgavara/goversine"
)

func main() {
	// New York decimal degrees (DD) coordinates are:
	// latitude 40.730610 N,
	// longitude 73.935242 W
	newYork := goversine.MustNewDDPoint(40.73061, -73.935242)

	// Madrid decimal degrees (DD) coordinates are:
	// latitude 40.416775 N,
	// longitude 3.703790 W
	madrid := goversine.MustNewDDPoint(40.416775, -3.70379)

	haversine := goversine.NewHaversine(goversine.Kilometre)
	distance := haversine.GetDistance(newYork, madrid)

	fmt.Printf("The distance from New York to Madrid is %.2f kilometres.\n", distance)
}
