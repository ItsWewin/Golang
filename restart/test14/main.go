package main

import "fmt"

func main() {
	sum := test(1, 2, 3, 4, 5)
	fmt.Printf("sum: %d\n", sum)
}

func test(n1 int, args ...int) int {
	sum := n1
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}

	return sum
}
