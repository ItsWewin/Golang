package main

import "fmt"

func main() {
	// 字符常量是使用但引号扩起来的字符，如下
	var c1 byte = 'x'

	// 字符常量直接输出，输出的是对应的码值
	fmt.Println(c1)

	// 使用 %c 输出字符常量本身
	fmt.Printf("c1 = %c\n", c1)

	// 一个整数，可以输出这个整数为码值，对饮的字符
	var c2 int = 120
	fmt.Printf("c2 = %c\n", c2)

	var c3 int = '深'
	fmt.Printf("c3 = %c, 对应的码值 %d\n", c3, c3)

	var val int = 28145
	fmt.Printf("%d 对应的字符是 %c\n", val, val)

	// 字符可参与进行运算，使用的是字符对应的
	var n1 = 10 + 'a'
	fmt.Printf("n1 = %v\n", n1)
}
