package main

import "fmt"

type A struct {
	Name string
	age  int
}

type B struct {
	Name  string
	score float64
}

type Uni struct {
	*A
	*B
	int
	n int
}

func main() {
	un := Uni{&A{"aName", 20}, &B{"bName", 99.4}, 12, 15}

	fmt.Println(un) //{{name 12} {99.99}}
	// fmt.Println("name: ", un.Name) // 因为 A、B 都有 Name 属性，但是 Uni 自身没有，所以这种导入会报错，必须要指明匿名结构体
	fmt.Println(un.A.Name) // 会访问到 A 的 Name，因为是有名导入，所以必须指定导入名称
	// fmt.Println(un.A.Name) // 有名导入这种写法错误

	fmt.Println(un.B.Name) // 会访问到 A 的 Name

	un.int = 13
	fmt.Println(un.int)
	fmt.Println(un.n)
}
