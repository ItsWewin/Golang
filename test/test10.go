package main

import (
	"fmt"
)

type Count int
type StringMap map[string]string
type FloatChan chan float64

func main() {
	var i Count = 7
	i++
	fmt.Println(i)

	sm := make(StringMap)
	sm["key1"] = "value1"
	sm["key2"] = "value2"
	fmt.Println(sm)

	fc := make(FloatChan, 1)
	fc <- 3.1415926
	fmt.Println(<-fc)
}