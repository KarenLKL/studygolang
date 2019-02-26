package main

import "fmt"

type treeNode struct {
	Value       int
	Left, Right *treeNode
}

func (node treeNode) print() {
	fmt.Println(node.Value)
}

func (node *treeNode) loop()  {
	if node==nil {
		return
	}
	node.Left.loop()
	node.print()
	node.Right.loop()
}

func main() {
	var root treeNode
	fmt.Println(root)
	root.Value = 3
	root.Left = &treeNode{Value: 2}
	root.Right = &treeNode{7, &treeNode{}, &treeNode{Value: 3}}

	root.print()
	root.loop()
}
