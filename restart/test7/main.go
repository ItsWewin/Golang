package main

import "fmt"

func main() {
	var x interface{}
	var y = 10

	x = y
	switch i := x.(type) {
	case nil:
		fmt.Printf("nil")
		fmt.Printf("%v", i)
	case int:
		fmt.Printf("int")
	default:
		fmt.Printf("others")
	}
}
