package main

import "fmt"

type Person struct {
	Name string
}

func (p Person) printName() {
	fmt.Printf("name: %v\n", p.Name)
	p.Name = "jack"
}

func (p *Person) setAndPrintName() {
	fmt.Printf("name: %v\n", p.Name)
	p.Name = "jack"
}

func main() {
	p := Person{"tom"}

	p.printName()    //name: tom
	(&p).printName() //name: tom

	p.setAndPrintName()    //name: tom
	(&p).setAndPrintName() //name: jack
}
