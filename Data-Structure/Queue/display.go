package main

import "log"

func (q *Queue) Display() {
	log.Println("Display Queue")
	node := q.left
	for i := 0; i < q.size; i++ {
		log.Println("Value : ", node.value)
		node = node.right
	}
}
