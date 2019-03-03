package main

import (
	"fmt"
)

type Point struct {x, y, z int}

func (point Point) String() string {
	return fmt.Sprintf("(%d, %d, %d)". print.x, print.y, print.z)
}

func main() {
	// massForPlanet := make(map[string]float64)
	// massForPlanet := map[string]float64{}
	// massForPlanet["key1"] = 0.01
	// massForPlanet["key2"] = 0.02
	// massForPlanet["key3"] = 0.03
	// massForPlanet["key4"] = 0.04
	// massForPlanet["key5"] = 0.05
	// fmt.Print(massForPlanet, "\n")

	triangle := make(map[*Point]string, 3)
	triangle[&Point{89, 47, 27}] = '1'
	triangle[&Point{11, 23, 56}] = '2'
	triangle[&Point{34, 57, 68}] = '3'
	fmt.Println(triangle)
}