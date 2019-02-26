package main

import "fmt"

// 使用chan 遍历树
type Node struct {
	Left  *Node
	Right *Node
	Value int
}

type TraverseFunc func(node *Node)

func (node *Node) TraverseWithChan() chan *Node {
	chanel := make(chan *Node)
	go func() {
		// 最后关闭chanel
		defer close(chanel)
		node.TraverseTree(func(node *Node) {
			fmt.Printf("Node value: %d \n", node.Value)
			chanel <- node
		})
	}()
	return chanel
}

// 遍历树，优先遍历左子树
func (node *Node) TraverseTree(f TraverseFunc) {
	if node == nil {
		return
	}
	node.Left.TraverseTree(f)
	f(node)
	node.Right.TraverseTree(f)
}

func main() {
	var root Node
	fmt.Println(root)
	root.Value = 3
	root.Left = &Node{Value: 2}
	root.Right = &Node{&Node{}, &Node{Value: 3}, 7}

	withChan := root.TraverseWithChan()
	maxValue := 0
	for c := range withChan {
		if c.Value > maxValue {
			maxValue = c.Value
		}
	}

	fmt.Printf("tree node maxValue:%d \n", maxValue)
}
