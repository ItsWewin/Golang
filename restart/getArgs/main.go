package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("args len: %v\n", len(os.Args))

	for i, v := range os.Args {
		fmt.Printf("index: %v, arg: %v\n", i, v)
	}
}
