package main

import (
	"fmt"
)

type Node struct {
	Val  int
	Next *Node
	//Prev *Node
}

type LL struct {
	Root *Node
}

func (l *LL) Insert(val int) {
	newNode := &Node{Val: val}
	if l.Root == nil {
		l.Root = newNode
		return
	}
	curr := l.Root

	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = newNode
	//newNode.Prev = curr
}

func (l *LL) Print() {
	if l.Root == nil {
		return
	}
	curr := l.Root
	for curr != nil {
		fmt.Printf(">-%d", curr.Val)
		curr = curr.Next
	}
	fmt.Println()

}

func removeNthFromEnd(head *Node, n int) *Node {
	dummy := &Node{Next: head}
	first, second := dummy, dummy
	for i := 0; i < n; i++ {
		first = first.Next
	}
	for first.Next != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
}

func (l *LL) ReverseLink() {
	var prev *Node
	curr := l.Root

	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}
	l.Root = prev
}

func main() {
	link := &LL{}
	link.Insert(20)
	link.Insert(30)
	link.Insert(40)
	link.Insert(50)
	fmt.Println("Before removing the nthNode")
	link.Print()

	// removeNthFromEnd(link.Root, 2)
	// link.Print()

	fmt.Println("After reverse linked list")
	link.ReverseLink()
	link.Print()

	removeNthFromEnd(link.Root, 2)
	link.Print()
}
