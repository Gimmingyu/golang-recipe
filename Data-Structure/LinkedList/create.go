package main

type LinkedList struct {
	size       uint
	headerNode Node // Dummy Node
}

type Node struct {
	value int
	next  *Node
}

func createNode(value int) *Node {
	return &Node{value: value, next: nil}
}

func createLinkedList() *LinkedList {
	return &LinkedList{}
}
