package main

import (
	"fmt"
)

func InsertStringSliceCopy(slice, insertion []string, index int) []string {
	result := make([]string, len(slice) + len(insertion))
	at := copy(result, slice[:index])
	at += copy(result[at:], insertion)
	copy(result[at:], slice[index:])

	return result
}

func main() {
	s := []string{"A", "B", "C", "D", "E", "F", "G"}
	t := []string{"K", "L", "M", "N"}
	u := []string{"m", "n", "o", "p", "q", "r"}

	s = append(s, "h", "i", "j")
	s = append(s, t...)
	s = append(s, u[2:5]...)
	fmt.Print(s, "\n")

	k := InsertStringSliceCopy(s, []string{"1", "2", "3"}, 0)
	fmt.Print(k, "\n")
}