package main

import (
	"errors"
	"fmt"
)

type Node struct {
	Key  int
	Val  int
	Next *Node
}

type HashTable struct {
	Table []*Node
	Size  int
}

func NewHasTable(size int) *HashTable {
	return &HashTable{Table: make([]*Node, size), Size: size}
}

func (h *HashTable) HashKey(key int) int {
	return key % h.Size
}

func (h *HashTable) Add(k, v int) {
	index := h.HashKey(k)
	newNode := &Node{Key: k, Val: v}
	if h.Table[index] != nil {
		newNode.Next = h.Table[index]
		h.Table[index] = newNode
	} else {
		h.Table[index] = newNode
	}
}

func (h *HashTable) Get(k int) (int, error) {
	index := h.HashKey(k)

	if h.Table[index] != nil {
		curr := h.Table[index]
		for curr != nil {
			if curr.Key == k {
				return curr.Val, nil
			}
			curr = curr.Next
		}
	}
	return 0, errors.New("not found")

}

func (h *HashTable) Remove(k int) bool {
	index := h.HashKey(k)
	if h.Table[index] != nil {
		var prev *Node
		curr := h.Table[index]
		for curr != nil {
			if curr.Key == k {
				if prev == nil {
					h.Table[index] = curr.Next
				} else {
					prev.Next = curr.Next
				}
				return true
			}
			prev = curr
			curr = curr.Next
		}
	}
	return false
}

func main() {
	hash := NewHasTable(5)
	hash.Add(1, 11)
	hash.Add(2, 22)
	hash.Add(3, 33)
	hash.Add(6, 66)
	hash.Add(7, 77)

	for key, val := range hash.Table {
		fmt.Println("key", key, "val", val)
	}

	fmt.Println(hash.Get(7))
	hash.Remove(6)

	for key, val := range hash.Table {
		fmt.Println("key", key, "val", val)
	}

}
