package main

import "fmt"

type Ainterface interface {
	say()
}

type Person struct{}

func (p Person) say() {
	fmt.Printf("say hello!\n")
}

func main() {
	var a Ainterface
	p := Person{}
	a = p
	a.say()
}
