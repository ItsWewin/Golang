package main

import (
	"fmt"
)

type composer struct {
	name string
	birthYear int
}

func main() {
	antonio := composer{"antonio Teixeira", 1708}
	fmt.Print(antonio.name, antonio.birthYear, "\n")

	antonio2 := new(composer)
	// antonio2.name, antonio2.birthYear = "name1", 1234
	antonio2.name = "name1"
	antonio2.birthYear = 1234
	fmt.Print(antonio2.name, antonio2.birthYear, "\n")

	article1 := &composer{}
	article1.name = "weiwei"
	article1.birthYear = 1999
	fmt.Print(article1.name, article1.birthYear, "\n")

	article2 := &composer{"wewin2", 1994}
	fmt.Print(article2.name, article2.birthYear)
}

