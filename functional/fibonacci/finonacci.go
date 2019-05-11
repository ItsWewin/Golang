package main

import "fmt"

func finonacci() func() int {
	a, b := 0, 1

	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	fb := finonacci()
	fmt.Println(fb())
	fmt.Println(fb())
	fmt.Println(fb())
	fmt.Println(fb())
	fmt.Println(fb())
	fmt.Println(fb())
	fmt.Println(fb())
	fmt.Println(fb())
}
