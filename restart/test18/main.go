package main

import "fmt"

func main() {
	f := AddUpper()
	num1 := f()
	num2 := f()

	fmt.Printf("num1: %d, num2: %d\n", num1, num2)
}

func AddUpper() func() int {
	var n = 0
	return func() int {
		n++
		return n
	}
}
