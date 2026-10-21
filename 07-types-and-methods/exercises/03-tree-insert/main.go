package main

import "fmt"

type Node struct {
	val         int
	left, right *Node
}

func (n *Node) Insert(v int) *Node {
	if n == nil {
		return &Node{val: v}
	}
	if v < n.val {
		n.left = n.left.Insert(v)
	} else if v > n.val {
		n.right = n.right.Insert(v)
	}
	return n
}

func (n *Node) Has(v int) bool {
	if n == nil {
		return false
	}
	if v < n.val {
		return n.left.Has(v)
	}
	if v > n.val {
		return n.right.Has(v)
	}
	return true
}

func main() {
	var root *Node
	root = root.Insert(5)
	root = root.Insert(3)
	root = root.Insert(10)
	fmt.Println(root.Has(3), root.Has(7))
}
