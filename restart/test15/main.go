package main

import "fmt"

var a = test()

func test() int {
	fmt.Printf("test()\n")
	return 90
}

func init() {
	fmt.Printf("init()\n")
}

func main() {
	fmt.Printf("main()\n")
}
