package main

func (q *Queue) Append(node *QueueNode) {
	if q.size == 0 {
		q.left = node
		q.right = node
		q.size++
		return
	}
	node.left = q.right
	node.right = nil
	q.right.right = node
	q.right = node
	q.size++
}

func (q *Queue) AppendLeft(node *QueueNode) {
	node.right = q.left
	node.left = nil
	q.left.left = node
	q.left = node
	q.size++
}
