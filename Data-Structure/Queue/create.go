package main

type IQueue interface {
	Append(node *QueueNode)
	AppendLeft(node *QueueNode)
	Pop() *QueueNode
	PopLeft() *QueueNode
	Display()
}

type Queue struct {
	size  int
	left  *QueueNode
	right *QueueNode
}

type QueueNode struct {
	value int
	left  *QueueNode
	right *QueueNode
}

func NewQueue() IQueue {
	return &Queue{}
}

func NewQueueNode(value int) *QueueNode {
	return &QueueNode{value: value}
}
