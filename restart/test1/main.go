package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// var i int8
	// var j int16
	// fmt.Println("i == j", i == j)

	var a = 23324124
	fmt.Printf("a 的类型是 %T, 范围是 %d", a, unsafe.Sizeof(a))
}
