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
