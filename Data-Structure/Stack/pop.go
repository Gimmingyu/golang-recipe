package main

func (s *Stack) popStackNode() *StackNode {
	if s.size <= 0 {
		panic("Empty stack")
	}
	top := s.topNode
	s.topNode = s.topNode.next
	top.next = nil
	s.size--
	return top
}
