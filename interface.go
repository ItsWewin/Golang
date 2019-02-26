package main

import (
	"fmt"
)

func main() {
	var i interface{} = 99
	var s interface{} = []string{"left", "right"}

	j := i.(int)
	fmt.Printf("%T -> %d\n", j, j)

	if i, ok := i.(int); ok {
		fmt.Printf("%T -> %d\n", i, i)
	}

	if s, ok := s.([]string); ok {
		fmt.Printf("%T -> %q\n", s, s)
	}
}