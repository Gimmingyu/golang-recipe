package main

func (q *Queue) Pop() *QueueNode {
	popNode := q.right
	q.right = q.right.left
	q.size--
	return popNode
}

func (q *Queue) PopLeft() *QueueNode {
	popNode := q.left
	q.left = q.left.right
	q.size--
	return popNode
}
