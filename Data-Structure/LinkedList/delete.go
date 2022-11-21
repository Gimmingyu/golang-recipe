package main

import (
	"fmt"
	"log"
)

func deleteNode(lst *LinkedList, target int) error {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered panic while inserting node: ", r)
		}
	}()
	if lst.size < 1 {
		return fmt.Errorf("Empty list\n")
	}
	if target == 0 {
		delNode := lst.headerNode.next
		lst.headerNode.next = delNode.next
		return nil
	}
	current := lst.headerNode.next
	for i := 0; i < target-1; i++ {
		current = current.next
	}
	current.next = current.next.next
	return nil
}
