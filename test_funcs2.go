package main

import (
	"fmt"
	"math"
)

func Heron(a, b, c int) float64 {
	x, y, z := float64(a), float64(b), float64(c)
	k := (x + y + z) / 2
	return math.Sqrt(k * (k - x) * (k - y) * (k - z))
}

func PythagoreanTriple(m, n int) (a, b, c int) {
	if m < n {
		m, n = n, m
	}

	return (m * n) - (n * n), (2 * m * n), (m * m) + (n * n)
}

func main() {
	article := Heron(PythagoreanTriple(1, 1+1))
	fmt.Print(article, "\n")
}