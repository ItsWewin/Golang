package main

import "fmt"

func main() {
	sum1, sum2 := test(1, 2)
	fmt.Printf("sum1: %d, sum2: %d\n", sum1, sum2)
}

func test(val1, val2 int) (sum1, sum2 int) {
	sum1 = val1 + val2
	sum2 = 2 * (sum1)
	return
}
