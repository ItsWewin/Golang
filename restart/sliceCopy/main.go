package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, 10)
	copy(slice2, slice1) // [1 2 3 4 5 0 0 0 0 0]
	fmt.Printf("slice2: %v\n", slice2)

	slice3 := make([]int, 3)
	copy(slice3, slice1)
	fmt.Printf("slice3: %v\n", slice3) //[1 2 3]

	str := "abcdefghijklmn"
	str2 := str[2:4]
	fmt.Printf("str2: %v\n", str2) // cd

	str3 := "abcdefghijklmnxxx"
	str4 := []byte(str3)
	str4[0] = 'z'
	str3 = string(str4)
	fmt.Printf("str3: %v\n", str3) // cd

}
