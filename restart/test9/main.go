package main

import "fmt"

func main() {
lable2:
	for i := 0; i < 10; i++ {
		// lable1:
		for j := 0; j < 5; j++ {
			fmt.Printf("i = %d, j = %d \n", i, j)
			if j == 2 {
				break lable2
			}
		}
	}
}
