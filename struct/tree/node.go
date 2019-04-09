package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

func main() {
	var root treeNode
	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}

	root.right.left = new(treeNode) // new(treeNode) 返回的就是地址
	root.right.left.value = 4
	root.right.left.left = nil
	root.right.left.right = nil

	fmt.Printf("%v, %T\n", root, root)
	fmt.Printf("%v, %T\n", root.left, root.left)
	fmt.Printf("%v, %T\n", root.right, root.right)
	fmt.Printf("%v, %T\n", root.right.left, root.right.left)

	root2 := treeNode{5, nil, nil}
	fmt.Printf("%v, %T\n", root2, root2)

	nodes := []treeNode{
		{5, nil, nil},
		{},
		{value: 4},
	}
	fmt.Printf("%v, %T", nodes, nodes)
}
