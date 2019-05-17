package main

import "fmt"

func main() {
	a := test
	fmt.Printf("address of a: %v\n", &a)
	fmt.Printf("address of test: %v", &test)
}

func test() {
	fmt.Printf("hello world\n")
}
