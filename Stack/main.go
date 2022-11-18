package main

import "log"

func main() {

	stack := NewStack()

	stack.insertNode(NewStackNode(1))
	stack.insertNode(NewStackNode(2))
	stack.insertNode(NewStackNode(3))
	stack.insertNode(NewStackNode(4))
	stack.insertNode(NewStackNode(5))

	stack.displayStack()
	log.Println("============================================================")

	log.Println("POP NODE = ", stack.popStackNode().value)

	log.Println("============================================================")
	log.Println("POP NODE = ", stack.popStackNode().value)

	log.Println("============================================================")
	log.Println("POP NODE = ", stack.popStackNode().value)

	log.Println("============================================================")
	log.Println("POP NODE = ", stack.popStackNode().value)

	log.Println("============================================================")
	log.Println("POP NODE = ", stack.popStackNode().value)
	log.Println("============================================================")
}
