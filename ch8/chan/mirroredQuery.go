package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(mirroredQuery())
}

func mirroredQuery() string {
	responses := make(chan string, 3)
	go func() {
		responses <- request("a.gopl.io")
	}()

	go func() {
		responses <- request("b.gopl.io")
	}()

	go func() {
		time.Sleep(1 * time.Minute)
		responses <- request("c.gopl.io")
	}()

	return <-responses
}

func request(hostname string) string {
	return hostname
}
