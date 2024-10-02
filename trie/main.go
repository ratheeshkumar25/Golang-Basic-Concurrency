package main

import "fmt"

const AlaphaSize = 26

type Node struct {
	Children [AlaphaSize]*Node
	IsEnd    bool
}

type Trie struct {
	Root *Node
}

func InitTrie() *Trie {
	result := &Trie{Root: &Node{}}
	return result
}

func (t *Trie) Insert(s string) {
	n := len(s)
	currNode := t.Root

	for i := 0; i < n; i++ {
		charIndex := s[i] - 'a'
		if currNode.Children[charIndex] == nil {
			currNode.Children[charIndex] = &Node{}
		}
		currNode = currNode.Children[charIndex]
	}
	currNode.IsEnd = true
}

func (t *Trie) Search(s string) bool {
	n := len(s)
	currNode := t.Root
	for i := 0; i < n; i++ {
		charIndex := s[i] - 'a'
		if currNode.Children[charIndex] == nil {
			return false
		}
		currNode = currNode.Children[charIndex]
	}

	if currNode.IsEnd == true {
		return true
	}

	return false
}

func (t *Trie) Print() {
	var printHelper func(node *Node, prefix string)

	printHelper = func(node *Node, prefix string) {
		if node.IsEnd {
			fmt.Println(prefix) // Print the word when we reach an end node
		}

		for i := 0; i < AlaphaSize; i++ {
			if node.Children[i] != nil { // If there is a child node
				printHelper(node.Children[i], prefix+string('a'+i)) // Append current character to prefix
			}
		}
	}

	printHelper(t.Root, "") // Start from root with an empty prefix
}

func main() {
	tries := InitTrie()
	toAdd := []string{
		"ratheesh",
		"vinoj",
		"musadehik",
	}

	for _, v := range toAdd {
		tries.Insert(v)
	}

	fmt.Println("Words available in trie:")
	tries.Print()

	// fmt.Println("trie", toAdd)

	fmt.Println("Word available in trie", tries.Search("vinoj"))
}
