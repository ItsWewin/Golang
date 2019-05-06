package main

import "fmt"

func main() {
	var root tree.treeNode
	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.left.right = createNode(2)

	root.right = &treeNode{5, nil, nil}

	root.right.left = new(treeNode) // new(treeNode) 返回的就是地址
	root.right.left.value = 4
	root.right.left.left = nil
	root.right.left.right = nil

	root.traversebyValue()
	fmt.Println("--------------------")
	root.traversebyPtr()
}
