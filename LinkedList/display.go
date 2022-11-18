package main

import (
	"fmt"
	"log"
)

func displayList(lst *LinkedList) error {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered panic while inserting node: ", r)
		}
	}()
	if lst.size < 1 {
		return fmt.Errorf("Empty list\n")
	}
	current := lst.headerNode.next
	index := 0
	for {
		log.Println("Display LinkedList")
		log.Printf("Index : %d\n", index)
		log.Printf("Value : %v\n", current.value)
		index++
		if current.next == nil {
			break
		}
		current = current.next
	}
	return nil
}
