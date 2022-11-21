package main

import "fmt"

type Node struct {
	Key   int
	left  *Node
	right *Node
}

func (n *Node) Insert(k int) {
	if n.Key < k {
		// move right
		if n.right == nil {
			n.right = &Node{Key: k}
		} else {
			n.right.Insert(k)
		}
	} else if n.Key > k {
		// move left
		if n.left == nil {
			n.left = &Node{Key: k}
		} else {
			n.left.Insert(k)
		}
	}
}

func (n *Node) Search(k int) bool {
	if n == nil {
		return false
	}
	if n.Key < k {
		// move right
		return n.right.Search(k)
	} else if n.Key > k {
		// move left
		return n.left.Search(k)

	}
	return true
}

func main() {
	tree := &Node{Key: 100}
	tree.Insert(52)
	tree.Insert(203)
	tree.Insert(19)
	tree.Insert(76)
	tree.Insert(150)
	tree.Insert(310)
	tree.Insert(7)
	tree.Insert(24)
	tree.Insert(88)
	tree.Insert(276)

	fmt.Println(tree.Search(76))
	fmt.Println(tree.Search(50))
	fmt.Println(tree.Search(100))
	fmt.Println(tree.Search(52))
}
