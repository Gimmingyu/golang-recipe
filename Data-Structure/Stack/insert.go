package main

func (s *Stack) insertNode(node *StackNode) {
	node.next = s.topNode
	s.topNode = node
	s.size++
}
