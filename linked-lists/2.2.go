package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type LinkedList struct {
	Head *Node
}

type Node struct {
	Value int
	Next  *Node
}

func kthToLast(ll LinkedList, k int) (*Node, error) {
	newLl := LinkedList{ll.Head}
	curr := ll.Head.Next
	currNewLl := newLl.Head
	for i := 1; i < k; i++ {
		if curr == nil {
			return nil, errors.New("Linked list has less than k nodes")
		}
		currNewLl = curr
		if i != k-1 {
			currNewLl = currNewLl.Next
		}
		curr = curr.Next
	}

	for curr != nil {
		newLl.Head = newLl.Head.Next
		curr = curr.Next
	}
	return newLl.Head, nil
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func main() {
	input := LinkedList{Head: &Node{1, &Node{2, &Node{3, &Node{4, &Node{1, &Node{2, &Node{2, &Node{3, &Node{5, nil}}}}}}}}}}
	kthNode, _ := kthToLast(input, 1)
	fmt.Println(kthNode)
}
