package main

import "log"

func (s *Stack) displayStack() {
	top := s.topNode
	for i := 0; i < s.size; i++ {
		log.Println("Value :", top.value)
		top = top.next
	}
}
