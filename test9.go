package main

import (
	"fmt"
)

func main() {
	var buffer [20]int
	fmt.Println(buffer)

	var buffer2 [3][3]int
	buffer2[1][0], buffer2[1][1], buffer2[1][2] = 2, 4, 6
	fmt.Println(buffer2)

	cities := [...]string{"shanghai", "Mumbai", "Istanbul"}
	fmt.Println(cities)
}