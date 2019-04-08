package main

import "fmt"

func passValue(arr [5]int) {
	for i, v := range arr {
		arr[i] = v + 10
	}
}

func passRef(arr *[5]int) {
	for i, v := range arr {
		arr[i] = v + 10
	}
}

func passSlic(arr []int) {
	for i, v := range arr {
		arr[i] = v + 100
	}
}

func main() {
	var arr1 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr1)

	passValue(arr1)
	fmt.Println(arr1)

	passRef(&arr1)
	fmt.Println(arr1)

	passSlic(arr1[3:])
	fmt.Println(arr1)
}
