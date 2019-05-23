package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int = 12
	rType := reflect.TypeOf(a)
	fmt.Println("rType", rType)

	rVal := reflect.ValueOf(a)
	fmt.Printf("rVal: %v\n", rVal)
	fmt.Println("rVal's kind", rVal.Kind())

	// reflect.ValueOf() 取的结果是一个 reflect.Value 类型，而不是 int 类型, 所以下面操作会报错
	// invalid operation: rVal + 2 (mismatched types reflect.Value and int)
	// newVal := rVal + 2

	// 这样写没有问题
	newVal := rVal.Int() + 2 // 类型必须正确，比如此处如果是 rVal.Float(), 这里肯定报错
	fmt.Printf("newVal: %d\n", newVal)

	// 转为 interface 类型
	rValToInter := rVal.Interface()
	fmt.Printf("rValToInter,type: %T, value: %v", rValToInter, rValToInter)

	// 使用类型断言，将 interface 类型转为 int 类型
	rValToInt := rValToInter.(int)
	newVal2 := rValToInt + 1
	fmt.Printf("newVal2, type: %T, value: %v", newVal2, newVal2)
}
