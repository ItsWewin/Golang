package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p1 := Person{"tom", 20}
	var p2 Person = Person{"tom", 20}
	var p3 = Person{"tom", 30}

	p4 := Person{
		Name: "tom",
		Age:  20,
	}

	var p5 = Person{
		Name: "tom",
		Age:  20,
	}

	var p6 Person = Person{
		Name: "tom",
		Age:  20,
	}

	fmt.Println(p1, p2, p3, p4, p5, p6) //{tom 20} {tom 20} {tom 30} {tom 20} {tom 20} {tom 20}
}
