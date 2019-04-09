package main

import "learngo/tree"

func main() {
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Left.Right = tree.CreateNode(2)

	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node) // new(Node) 返回的就是地址
	root.Right.Left.Value = 4
	root.Right.Left.Left = nil
	root.Right.Left.Right = nil

	root.TraversebyValue()
	root.TraversebyPtr()
	root.SetValue(10)
}
