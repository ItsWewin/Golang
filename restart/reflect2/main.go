package main

import (
	"fmt"
	"reflect"
)

func reflectTest(value interface{}) {
	rVal := reflect.ValueOf(value)
	fmt.Println("rVal", rVal)

	rValToInter := rVal.Interface()
	fmt.Printf("type of rValToInter: %T, value of rValToInter: %v\n", rValToInter, rValToInter)

	value2 := rValToInter.(float64)
	num := value2 + 12.3
	fmt.Printf("num: %v\n", num)
}

func main() {
	var v float64 = 1.2
	reflectTest(v)
}
