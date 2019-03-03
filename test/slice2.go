package main

import (
	"fmt"
)

func main() {
	buffer := make([]byte, 20, 60)
	grid1 := make([][]int, 3)
	fmt.Print(buffer, "\n")
	fmt.Print(grid1, "\n")
	for i := range grid1 {
		grid1[i] = make([]int, 3)
	}

	grid1[1][0], grid1[1][1], grid1[1][2] = 8, 6, 4
	fmt.Print(grid1, "\n")

	grid2 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Print(grid2, "\n")

	cities := make([][]int, 4)
	fmt.Print(cities)

	fmt.Print("================", "\n")

	s := []string{"a", "b", "c", "d", "e", "f", "g"}
	t := s[2:6]
	fmt.Println(t, s, "=", s[:4], "+", s[4:])
	s[3] = "x"
	fmt.Println(t, "|", s)

	amounts := []float64{123.12, 456.23, 124.55, 555.11, 2, 3, 2.1}
	sum := 0.0
	for _, amount := range amounts[:3] {
		sum += amount
	}

	sum2 := 0.0
	for _, amount := range amounts {
		sum2 += amount
	}
	fmt.Print(sum2, "\n")
}