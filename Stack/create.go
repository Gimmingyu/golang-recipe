package main

type IStack interface {
	insertNode(node *StackNode)
	displayStack()
	popStackNode() *StackNode
}

type Stack struct {
	size       int
	topNode    *StackNode
	bottomNode *StackNode
}

func NewStack() IStack {
	return &Stack{}
}

type StackNode struct {
	value int
	next  *StackNode
}

func NewStackNode(value int) *StackNode {
	return &StackNode{value: value}
}
