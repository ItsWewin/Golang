package main

import "fmt"

type Person struct {
	Name string
}

func (p Person) printName() {
	fmt.Printf("name is %v\n", p.Name)
}

func (p Person) saySome(word string) {
	fmt.Printf("%v saied %v\n", p.Name, word)
}

func main() {
	p := Person{"tom"}
	p.printName()
	p.saySome("hello")
}
