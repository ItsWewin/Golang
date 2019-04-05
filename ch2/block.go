package main

import (
	"fmt"
)

func main() {
	if x := 1; x == 0 {
		fmt.Println(x)
	} else if y := 2; x == y {
		fmt.Println(x)
	} else {
		fmt.Println(x, y)
	}
}
