package main

import (
	"bufio"
	"fmt"
	"os"
)

// exit input
// ctrl + d
// cat inputFile | go run dup1.go
func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		if 1 <= n {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}