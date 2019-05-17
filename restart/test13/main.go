package main

import "fmt"

type myFun func(int, int) int

func main() {
	val := myFun1(getSum, 1, 3)
	fmt.Printf("val: %d\n", val)
}

func myFun1(funvar myFun, num1, num2 int) int {
	return funvar(num1, num2)
}

func getSum(num1, num2 int) int {
	return num1 + num2
}
