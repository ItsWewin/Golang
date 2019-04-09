package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

func createNode(value int) *treeNode {
	return &treeNode{value: value} // 能返回局部变量的地址
}

func main() {
	var root treeNode
	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.left.right = createNode(2)

	root.right = &treeNode{5, nil, nil}

	root.right.left = new(treeNode) // new(treeNode) 返回的就是地址
	root.right.left.value = 4
	root.right.left.left = nil
	root.right.left.right = nil

	fmt.Printf("%v, %T, %d\n", root, root, root)
	fmt.Printf("%v, %T\n", root.left, root.left)
	fmt.Printf("%v, %T\n", root.right, root.right)
	fmt.Printf("%v, %T\n", root.right.left, root.right.left)

	root2 := treeNode{5, nil, nil}
	fmt.Printf("%v, %T\n", root2, root2)

	nodes := []treeNode{
		*createNode(5),
		{},
		{value: 4},
	}
	fmt.Printf("%v, %T", nodes, nodes)
}
