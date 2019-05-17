package main

import "fmt"

var (
	mySum = func(n1 int, n2 int) int {
		return n1 + n2
	}
)

func init() {
	sum := mySum(1, 2)
	fmt.Printf("sum: %d\n", sum)
}

func main() {
	sum := mySum(1, 2)
	fmt.Printf("sum: %d\n", sum)
}
