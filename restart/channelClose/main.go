package main

import "fmt"

func main() {
	intChan := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan <- i
	}

	close(intChan)
	for value := range intChan {
		fmt.Printf("value: %v\n", value)
	}
}
