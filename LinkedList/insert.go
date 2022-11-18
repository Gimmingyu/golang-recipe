package main

import "log"

func insertNode(lst *LinkedList, node *Node) error {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered panic while inserting node: ", r)
		}
	}()
	if lst.size == 0 {
		lst.headerNode.next = node
		lst.size++
		return nil
	}
	current := lst.headerNode.next
	for {
		if current.next == nil {
			break
		}
		current = current.next
	}
	current.next = node
	lst.size++
	return nil
}
