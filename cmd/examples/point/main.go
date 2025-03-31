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
