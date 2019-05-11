package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	fmt.Println(len(ch))
	fmt.Println(cap(ch))
	ch <- 1
	fmt.Println(len(ch))
	fmt.Println(cap(ch))
	ch <- 2
	fmt.Println(len(ch))
	fmt.Println(cap(ch))
	ch <- 3
	fmt.Println(len(ch))
	fmt.Println(cap(ch))
	// ch <- 4
}
