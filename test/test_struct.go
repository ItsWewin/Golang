package main

import (
	"fmt"
)

func main() {
	points := [][2]int{{4, 6}, {}, {-7, 11}, {15, 17}, {14, -8}}
	for _, point := range points {
		fmt.Printf("(%d, %d)\n", point[0], point[1])
	}
}