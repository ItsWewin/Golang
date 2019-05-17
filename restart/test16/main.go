package main

import "fmt"

func main() {
	res1 := func(n1 int, n2 int) int {
		return n1 + n2
	}(1, 2)

	fmt.Printf("res1= %d\n", res1)

	a := func(n1 int, n2 int) int {
		return n1 - n2
	}

	res2 := a(10, 30)
	fmt.Printf("res2: %d\n", res2)
}
