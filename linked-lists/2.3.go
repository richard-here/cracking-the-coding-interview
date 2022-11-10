package main

import (
	"encoding/json"
	"fmt"
)

type LinkedList struct {
	Head *Node
}

type Node struct {
	Value int
	Next  *Node
}

func deleteMiddleNode(ll LinkedList, n *Node) LinkedList {
	curr := ll.Head
	var prev *Node
	for curr != nil {
		if curr == n {
			prev.Next = curr.Next
			break
		}
		prev = curr
		curr = curr.Next
	}
	return ll
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func main() {
	nodeToRemove := &Node{1, &Node{2, &Node{2, &Node{3, &Node{5, nil}}}}}
	input := LinkedList{Head: &Node{1, &Node{2, &Node{3, &Node{4, nodeToRemove}}}}}
	fmt.Println(prettyPrint(deleteMiddleNode(input, nodeToRemove)))
}
