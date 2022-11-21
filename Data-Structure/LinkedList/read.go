package main

func getNode(lst *LinkedList, index uint) *Node {
	if lst.size == 0 {
		return nil
	}
	var i uint
	current := lst.headerNode.next
	for i = 0; i < lst.size && i < index; i++ {
		current = current.next
	}
	return current
}
