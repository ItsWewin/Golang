package main

import "fmt"

func main() {
	var x interface{}
	var f float64 = 43.2
	x = f
	fmt.Printf("f type %T\n", f)
	fmt.Printf("x type %T\n", x)

	// f = x  // 这种写法会出错
	f = x.(float64) // 使用类型断言，断言失败会报

	v, ok := x.(float32) // 使用这种方法可以，可以保证断言失败后不抛异常
	if ok {
		fmt.Printf("assertion success: %v", v)
	} else {
		fmt.Printf("assertion error")
	}
}
