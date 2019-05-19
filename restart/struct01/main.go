package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	person := Person{"name", 20}
	value, _ := json.Marshal(person)
	str := string(value)
	fmt.Printf("value: %v", str)
}
