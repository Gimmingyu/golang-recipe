package main

import (
	"log"
)

func main() {
	lst := createLinkedList()

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
