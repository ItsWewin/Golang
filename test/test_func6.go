package main

import (
	"fmt"
)

func Minimum(first interface{}, rest ...interface{}) interface{} {
	minimum := first
	for _, x := range rest {
		switch x := x.(type) {
		case int:
			if x < minimum.(int) {
				minimum = x
			}
		case float64:
			if x < minimum.(float64) {
				minimum =x
			}
		}
	}
	return minimum
}

func main() {
	i := Minimum(4, 3, 1, 9, 8).(int)
	fmt.Print(i, "\n")
	fmt.Print(123.(int))
}