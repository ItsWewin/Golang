package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// counter
	go func() {
		for x := 0; ; x++ {
			if x > 100000 {
				close(naturals)
				break
			}
			naturals <- x
		}
	}()

	go func() {
		for {
			x, ok := <-naturals
			fmt.Println(ok)
			if !ok {
				break
			}
			squares <- x * x
		}
		close(squares)
	}()

	for {
		squares, ok := <-squares
		if ok {
			fmt.Println(squares)
		}
	}
}
