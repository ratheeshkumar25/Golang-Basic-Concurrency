package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

type BS struct {
	Root *Node
}

func (b *BS) Insert(val int) {
	newNode := &Node{Val: val}
	if b.Root == nil {
		b.Root = newNode
		return
	}
	queue := []*Node{b.Root}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.Left == nil {
			curr.Left = newNode
			return
		} else {
			queue = append(queue, curr.Left)
		}
		if curr.Right == nil {
			curr.Right = newNode
			return
		} else {
			queue = append(queue, curr.Right)
		}
	}
}

func (b *BS) InOrder(node *Node) {
	if node != nil {
		b.InOrder(node.Left)
		fmt.Println(node.Val)
		b.InOrder(node.Right)

	}
}

func main() {
	bt := &BS{}
	bt.Insert(10)
	bt.Insert(20)
	bt.Insert(30)
	bt.Insert(50)

	bt.InOrder(bt.Root)
}
