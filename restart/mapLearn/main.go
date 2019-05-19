package main

import "fmt"

func main() {
	var myMap map[string]string
	// 没有分配空间，所以会报错
	// myMap["key1"] = "value"          // panic: assignment to entry in nil map
	fmt.Printf("myMap: %v\n", myMap) //map: map[]

	myMap = make(map[string]string, 10)
	myMap["key1"] = "value1"
	myMap["key2"] = "value2"
	fmt.Printf("myMap: %v\n", myMap) //myMap: map[key1:value1 key2:value2]

	myMap2 := make(map[string]string)
	myMap2["key1"] = "value1"
	myMap2["key2"] = "value2"
	fmt.Printf("myMap2: %v\n", myMap2)

	var myMap3 map[string]string = map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	fmt.Printf("myMap3: %v\n", myMap3)

	value2 := myMap3["key1"]
	fmt.Printf("value2: %v\n", value2)

	value2, fined := myMap3["key"]
	fmt.Printf("fined %v, value2: %v\n", fined, value2)
}
