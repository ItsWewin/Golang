package main

import "fmt"

func main() {
	num1 := new(int)
	fmt.Printf("num2 type: %T, num2 value: %v, num2 地址： %v\n", num1, num1, &num1)
	fmt.Printf("num1 指向的值为： %v\n", *num1)
	*num1 = 3
	fmt.Printf("num1 指向的值为： %v\n", *num1)
}
