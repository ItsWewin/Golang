package main

import (
	"fmt"
)

func inflate(numbers []int, factor int) {
	for i := range numbers {
		numbers[i] *= factor
	}
}

func main() {
	grades := []int{87, 55, 43, 71}
	inflate(grades, 3)
	fmt.Println(grades)
}