package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))
}

func passValue(s []int) {
	for i, _ := range s {
		s[i] += 100
	}
}

func arrPassValue(arr *[5]int) {
	for i, _ := range arr {
		arr[i] += 1000
	}
}

func main() {
	// defined slice
	var s1 = []int{1, 2, 3, 4, 5}
	printSlice(s1)
	passValue(s1)
	printSlice(s1)

	fmt.Println("-------------------")
	var arr1 = [...]int{1, 2, 3, 4, 5}
	printSlice(arr1[:])
	arrPassValue(&arr1)
	printSlice(arr1[:])

	fmt.Println("-------------------")
	s1 = s1[1:]
	printSlice(s1[:])

	var s2 = [][]int{}
	tem := []int{1, 2, 3}
	s3 := append(s2, tem)

	tem = []int{4, 5, 6}
	s3 = append(s3, tem)

	fmt.Println(s3)
}
