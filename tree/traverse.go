package tree

func (node Node) TraversebyValue() {
	if node.Left != nil {
		node.Left.TraversebyValue()
	}

	node.Print()

	if node.Right != nil {
		node.Right.TraversebyValue()
	}
}

func (node *Node) TraversebyPtr() {
	if node == nil {
		return
	}
	node.Left.TraversebyPtr()
	node.Print()
	node.Right.TraversebyPtr()
}
