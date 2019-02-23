package main

import (
	"fmt"
)

func main() {
	name1 := "we"
	name2 := "win"
	name := name1 + name2
	// str1 := "我是wewin"
	// str2 := "wewin"
	phrase := "val og tt"
	fmt.Printf("String: \"%s\"\n", phrase);
	for index, char := range phrase {
		fmt.Printf("%-2d    %U    '%c'    %X\n",
				index, char, char, []byte(string(char)))
	}

	fmt.Printf("name1: %s, name2: %s, name: %s\n", name1, name2, name)
	// fmt.Printf("string: %s\n", string(U+00C5))
}
