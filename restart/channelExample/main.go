package main

import (
	"fmt"
)

func writeInto(intChan chan int) {
	for i := 0; i < 50; i++ {
		intChan <- i
	}
	close(intChan)
}

func readFrom(intChan chan int, finished chan bool) {
	for {
		value, ok := <-intChan
		if !ok {
			break
		}
		fmt.Printf("value: %v\n", value)
	}

	finished <- true
	close(finished)
}

func main() {
	intChan := make(chan int, 1)
	finished := make(chan bool, 1)

	go writeInto(intChan)
	go readFrom(intChan, finished)

	for {
		_, ok := <-finished
		if !ok {
			break
		}
	}
}
