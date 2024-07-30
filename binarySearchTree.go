package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(val int) {
	if val < n.Value {
		// move left
		if n.Left == nil {
			n.Left = &Node{Value: val}
		} else {
			n.Left.Insert(val)
		}
	} else if val > n.Value {
		// move right
		if n.Right == nil {
			n.Right = &Node{Value: val}
		} else {
			n.Right.Insert(val)
		}
	}
}

func (n *Node) Print() {
	if n != nil {
		fmt.Println(n.Value)
		n.Left.Print()
		n.Right.Print()
	}
	return
}

func (n *Node) Search(val int) bool {
	if n == nil {
		return false
	}
	if val < n.Value {
		// search left
		return n.Left.Search(val)
	} else if val > n.Value {
		// search right
		return n.Right.Search(val)
	}
	return true
}

func main() {
	bst := &Node{Value: 50}

	bst.Insert(10)
	bst.Insert(100)
	bst.Insert(1)
	bst.Insert(5)

	bst.Print()

	fmt.Println(bst.Search(2))
}
