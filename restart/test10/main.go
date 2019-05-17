package main

import "fmt"

func main() {
lable1:
	for i := 0; i < 10; i++ {
		for j := 0; j < 5; j++ {
			if j == 3 {
				continue lable1
			}

			fmt.Printf("i = %d, j = %d\n", i, j)
		}
	}
}
