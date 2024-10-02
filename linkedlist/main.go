package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type LinkList struct {
	Head *Node
}

func (l *LinkList) Add(val int) {
	newNode := &Node{Val: val}

	if l.Head == nil {
		l.Head = newNode
	} else {
		current := l.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
}

func (l *LinkList) Display() {
	curr := l.Head

	for curr != nil {
		fmt.Printf("%d->", curr.Val)
		curr = curr.Next
	}
}

func (l *LinkList) ReverseLink() {
	var prev *Node
	current := l.Head

	for current != nil {
		temp := current.Next
		current.Next = prev
		prev = current
		current = temp
	}
	l.Head = prev
}

func main() {

	link := LinkList{}
	link.Add(20)
	link.Add(30)
	link.Add(40)
	link.Add(60)
	link.Add(80)
	link.Add(90)

	link.Display()

	link.ReverseLink()
	fmt.Println("after reverese the linked list")
	link.Display()

}
