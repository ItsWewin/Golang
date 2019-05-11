package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdout.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	select {
	case <-time.After(10 * time.Second):
	case <-abort:
		fmt.Println("Launch aborted!")
		return
	}

	launch()
}

func launch() {
	fmt.Println("launch it!")
}
