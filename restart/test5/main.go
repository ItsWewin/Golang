package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str1 = "123"
	var n1 int64
	n1, _ = strconv.ParseInt(str1, 10, 64)
	fmt.Printf("n1 = %d\n", n1)

	// 不能转成功的，会转成零值
	var str2 = "123hello123"
	var n2 int64
	n2, _ = strconv.ParseInt(str2, 10, 64)
	fmt.Printf("n2 = %d\n", n2)

	var str3 = "true"
	var b1 bool
	b1, _ = strconv.ParseBool(str3)
	fmt.Printf("b1 = %v\n", b1)

	var str4 = "false"
	var b2 bool
	b1, _ = strconv.ParseBool(str4)
	fmt.Printf("b2 = %v\n", b2)

	var str5 = "hello"
	var b3 bool
	b1, _ = strconv.ParseBool(str5)
	fmt.Printf("b2 = %v\n", b3)

	fmt.Printf("%b", str5)
}
