package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

type BT struct {
	Root *Node
}

func (b *BT) Insert(val int) {
	newNode := &Node{Val: val}
	if b.Root == nil {
		b.Root = newNode
		return
	}
	curr := b.Root

	for {
		if val < curr.Val {
			if curr.Left == nil {
				curr.Left = newNode
				return
			}
			curr = curr.Left
		} else {
			if curr.Right == nil {
				curr.Right = newNode
				return
			}
			curr = curr.Right
		}
	}

}

func (b *BT) Inorder(node *Node) {
	if node != nil {
		b.Inorder(node.Left)
		fmt.Println(node.Val)
		b.Inorder(node.Right)
	}
}

func nthSmall(node *Node, k int) int {
	c := make(chan int)
	go InOrderNth(node, c)
	for i := 0; i < k-1; i++ {
		<-c
	}
	return <-c

}

func InOrderNth(root *Node, c chan int) {
	if root == nil {
		return
	}
	InOrderNth(root.Left, c)
	c <- root.Val
	InOrderNth(root.Right, c)
}

func main() {

	bt := &BT{}
	bt.Insert(30)
	bt.Insert(20)
	bt.Insert(50)
	bt.Insert(60)
	bt.Inorder(bt.Root)

	b := nthSmall(bt.Root, 4)
	fmt.Println("Nth smallest ", b)

}
