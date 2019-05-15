package main

import "fmt"

func main() {
	var str = "abcd深圳"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%v ", str[i]) // 97 98 99 100 230 183 177 229 156 179
	}
	fmt.Println()
	for _, c := range str {
		// fmt.Printf("%v ", c) //97 98 99 100 28145 22323
		fmt.Printf("%c ", c) //a b c d 深 圳
	}

	fmt.Println()
	str2 := []rune(str)
	for i := 0; i < len(str2); i++ {
		// fmt.Printf("%v ", str2[i]) //97 98 99 100 28145 22323
		fmt.Printf("%c ", str2[i]) //a b c d 深 圳
	}
}
