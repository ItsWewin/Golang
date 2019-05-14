package main

import "fmt"

func main() {
	var a byte = 13
	var b uint8 = 13
	fmt.Println("a == b", a == b)

	var c rune = 13
	var d int32 = 13
	fmt.Println("c == d", c == d)
}
