package main

import "fmt"

func main() {
	test()

	fmt.Println("test test test test")
}

func test() {
	defer func() {
		// error := recover()
		// if error != nil {
		// 	fmt.Printf("some error: %v", error)
		// }

		if err := recover(); err != nil {
			fmt.Printf("some error: %v", error)
		}
	}()
	num1 := 1
	num2 := 0
	var num = num1 / num2
	fmt.Println(num)
}
