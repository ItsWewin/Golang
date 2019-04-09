package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func CreateNode(Value int) *Node {
	return &Node{Value: Value} // 能返回局部变量的地址
}

// 拷贝一个 Node 的副本
func (node Node) print() {
	fmt.Println(node.Value)
}

// 会智能的拿出 Node 的地址
// nil 指针也可以作为接受者
func (node *Node) SetValue(Value int) {
	if node == nil {
		fmt.Println("the tree is nil")
		return
	}
	node.Value = Value
}
