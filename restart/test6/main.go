package main

import "fmt"

func main() {
	if age := 20; age > 40 {
		fmt.Printf("You are so old\n")
	} else if name := "weiwei"; name == "weiwei1" {
		fmt.Printf("your age %d\n", age)
	} else {
		fmt.Printf("age: %d\n", age)
		fmt.Printf("name: %s\n", name)
	}
}
