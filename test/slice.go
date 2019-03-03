package main

import (
	"fmt"
)

func main() {
	// s := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s := new([7]int)[:]
	s[0] = 1
	s[2] = 2
	s[3] = 3
	s[4] = 4
	s[5] = 5
	s[6] = 6

	fmt.Print(len(s), "\n")
	fmt.Print(cap(s), "\n")

	fmt.Print(len(s[:5]), "\n")
	fmt.Print(cap(s[:5]), "\n")

	fmt.Print(len(s[2:5]), "\n")
	fmt.Print(cap(s[2:5]), "\n")

}


