package main

import "fmt"

type MaxHeap struct {
	array []int
}

func (m *MaxHeap) Insert(key int) {
	m.array = append(m.array, key)
	m.maxHeapifyUp(len(m.array) - 1)
}

func (m *MaxHeap) Extract() int {
	extracted := m.array[0]
	length := len(m.array) - 1
	if length == 0 {
		return -1
	}

	m.array[0] = m.array[length]
	m.array = m.array[:length]

	m.maxHeapifyDown(0)
	return extracted
}

func (m *MaxHeap) maxHeapifyUp(idx int) {
	for m.array[parent(idx)] < m.array[idx] {
		m.swap(parent(idx), idx)
		idx = parent(idx)
	}
}

func (m *MaxHeap) maxHeapifyDown(idx int) {
	lastIndex := len(m.array) - 1
	l, r := left(idx), right(idx)
	childToCompare := 0

	for l <= lastIndex {
		// when left child is only child
		if l == lastIndex {
			childToCompare = l
		} else if m.array[l] > m.array[r] { // when left child is larger
			childToCompare = l
		} else { // when right child is larger
			childToCompare = r
		}
		// compare array value of current index to larger child and swap if smaller
		if m.array[idx] < m.array[childToCompare] {
			m.swap(idx, childToCompare)
			idx = childToCompare
			l, r = left(idx), right(idx)
		} else {
			return
		}
	}
}

func parent(index int) int {
	return (index - 1) / 2
}

func left(index int) int {
	return (index * 2) + 1
}

func right(index int) int {
	return (index * 2) + 2
}

func (m *MaxHeap) swap(idx1, idx2 int) {
	m.array[idx1], m.array[idx2] = m.array[idx2], m.array[idx1]
}

func main() {
	m := &MaxHeap{}
	fmt.Println(m)

	buildHeap := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}

	for i := 0; i < 5; i++ {
		fmt.Println(m.Extract())
		fmt.Println(m)
	}
}
