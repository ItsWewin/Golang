package main

import (
	"fmt"
	"strings"
)

func IntForBool(b bool) int {
	if b {
		return 1
	}

	return 0
}

func main() {
	name1 := "we"
	name2 := "win"
	name := name1 + name2
	fmt.Printf("name1: %s, name2: %s, name: %s\n", name1, name2, name)

	phrase := "val og tt 魏巍"
	fmt.Printf("String: \"%s\"\n", phrase);
	for index, char := range phrase {
		fmt.Printf("%-2d    %U  %v  '%c'    %X\n",
				index, char, char, char, []byte(string(char)))
	}

	line := "xxx ddd ttt"
	i := strings.Index(line, " ")
	firstWord := line[0:]
	fmt.Printf("first word: %s\n", firstWord)
	fmt.Printf("first blank index: %d\n", i)

	// test %t
	fmt.Printf("%t %t\n", true, false)
	fmt.Printf("%d %d\n", IntForBool(true), IntForBool(false))

	// test print
	fmt.Printf("|%b|%9b|%-9b|%09b|% 9b\n", 37, 37, 37, 37, 37)
}
