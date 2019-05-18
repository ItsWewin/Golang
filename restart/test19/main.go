package main

import (
	"fmt"
	"strings"
)

func main() {
	str := strings.Trim("abd123456bac", "abc")
	fmt.Printf("str: %s\n", str)

	fmt.Print(strings.Trim("¡¡¡Hello, Gophers!!!", "!¡"))
}
