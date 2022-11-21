package main

import "log"

func main() {
	q := NewQueue()

	q.Append(NewQueueNode(1))
	q.Append(NewQueueNode(2))
	q.Append(NewQueueNode(3))
	q.Append(NewQueueNode(4))
	q.Append(NewQueueNode(5))
	q.Append(NewQueueNode(6))

	q.Display()
	log.Println("================================================================")

	log.Printf("Pop Node : %v\n", q.Pop().value)
	log.Printf("Pop Left Node : %v\n", q.PopLeft().value)
	q.Display()
	log.Println("================================================================")
	log.Printf("Pop Node : %v\n", q.Pop().value)
	log.Printf("Pop Left Node : %v\n", q.PopLeft().value)
	q.Display()
	log.Println("================================================================")
	log.Printf("Pop Node : %v\n", q.Pop().value)
	log.Printf("Pop Left Node : %v\n", q.PopLeft().value)
	q.Display()
	log.Println("================================================================")
}
