package main

import (
	"fmt"
	"log"
)

type LinkedList struct {
	size       uint
	headerNode Node // Dummy Node
}

type Node struct {
	value int
	next  *Node
}

func createNode(value int) *Node {
	return &Node{value: value, next: nil}
}

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

func main() {
	lst := &LinkedList{
		size:       0,
		headerNode: Node{},
	}

	// Create Node
	n1 := createNode(1)
	n2 := createNode(2)
	n3 := createNode(3)
	n4 := createNode(4)
	n5 := createNode(5)

	// Insert Node
	if err := insertNode(lst, n1); err != nil {
		log.Panicf("Error occured while inserting node : %v\n", err)
	}
	if err := insertNode(lst, n2); err != nil {
		log.Panicf("Error occured while inserting node : %v\n", err)
	}
	if err := insertNode(lst, n3); err != nil {
		log.Panicf("Error occured while inserting node : %v\n", err)
	}
	if err := insertNode(lst, n4); err != nil {
		log.Panicf("Error occured while inserting node : %v\n", err)
	}
	if err := insertNode(lst, n5); err != nil {
		log.Panicf("Error occured while inserting node : %v\n", err)
	}
	log.Print("================================================================")

	if err := displayList(lst); err != nil {
		log.Panicf("Error occured while displaying node : %v\n", err)
	}

	log.Print("================================================================")
	if err := deleteNode(lst, 3); err != nil {
		log.Panicf("Error occured while deleting node : %v\n", err)
	}
	if err := displayList(lst); err != nil {
		log.Panicf("Error occured while displaying node : %v\n", err)
	}
	log.Print("================================================================")

	if err := deleteNode(lst, 3); err != nil {
		log.Panicf("Error occured while deleting node : %v\n", err)
	}
	if err := displayList(lst); err != nil {
		log.Panicf("Error occured while displaying node : %v\n", err)
	}
	
}
