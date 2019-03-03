package main

import (
	"fmt"
)

func main() {
	i := 1
	j := 2
	i, j, product := swapAndProduct(i, j)
	fmt.Printf("i: %d, j: %d, product: %d\n", i, j, product)
}

func swapAndProduct(x, y int) ( int, int, int) {
	if x > y {
		x, y = y, x
	}

	return x, y,  x * y
}