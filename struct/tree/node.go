package tree

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

func createNode(value int) *treeNode {
	return &treeNode{value: value} // 能返回局部变量的地址
}

// 拷贝一个 treeNode 的副本
func (node treeNode) print() {
	fmt.Println(node.value)
}

// 会智能的拿出 treeNode 的地址
// nil 指针也可以作为接受者
func (node *treeNode) setValue(value int) {
	if node == nil {
		fmt.Println("the tree is nil")
		return
	}
	node.value = value
}

func (node treeNode) traversebyValue() {
	if node.left != nil {
		node.left.traversebyValue()
	}

	node.print()

	if node.right != nil {
		node.right.traversebyValue()
	}
}

func (node *treeNode) traversebyPtr() {
	if node == nil {
		return
	}
	node.left.traversebyPtr()
	node.print()
	node.right.traversebyPtr()
}
