package main

import "learngo/tree"

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrderTraversal() {
	if myNode == nil || myNode.node == nil {
		return
	}

	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}
	left.postOrderTraversal()
	right.postOrderTraversal()

	myNode.node.Print()
}

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

	myRoot := myTreeNode{&root}
	myRoot.postOrderTraversal()
}
