package main

import (
	"fmt"
	"log"
)

const ArraySize = 7

// HashTable
type HashTable struct {
	array [ArraySize]*Bucket
}

// Bucket
type Bucket struct {
	headerNode *BucketNode
}

// BucketNode
type BucketNode struct {
	key  string
	next *BucketNode
}

func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

func (b *Bucket) insert(key string) {
	if !b.search(key) {
		newNode := &BucketNode{key: key}
		newNode.next = b.headerNode
		b.headerNode = newNode
	} else {
		fmt.Println("Already exists")
	}
}

func (b *Bucket) search(key string) bool {
	currentNode := b.headerNode
	for currentNode != nil {
		if currentNode.key == key {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

func (b *Bucket) delete(key string) {
	prevNode := b.headerNode
	for prevNode.next != nil {
		if prevNode.next.key == key {
			prevNode.next = prevNode.next.next
			return
		}
		prevNode = prevNode.next
	}
}

// Hash
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &Bucket{}
	}
	return result
}

func main() {
	testHashTable := Init()
	fmt.Println(testHashTable)

	testHashTable.Insert("RANDY")
	fmt.Println(testHashTable.Search("RANDY"))
	testHashTable.Delete("RANDY")
	fmt.Println(testHashTable.Search("RANDY"))

	for _, v := range testHashTable.array {
		if v.headerNode != nil {
			log.Println(v.headerNode.key)
		}
	}
}
