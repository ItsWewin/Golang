package main

import "fmt"

type myInterface interface {
	method1()
	method2()
	method3()
}

type struct1 struct{}
type struct2 struct{}
type struct3 struct{}

type unic struct {
	struct1
	struct2
	struct3
}

func (s struct1) method1() {
	fmt.Println("method1")
}

func (s struct2) method2() {
	fmt.Println("methdo2")
}

func (s struct3) method3() {
	fmt.Println("methdo3")
}

func useing(s myInterface) {
	s.method1()
	s.method2()
	s.method3()
}

func main() {
	s := unic{}
	useing(s)
}
