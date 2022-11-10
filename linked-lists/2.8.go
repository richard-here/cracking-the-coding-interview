package main

import (
	"fmt"
)

type LinkedList struct {
	Head *Node
}

type Node struct {
	Value int
	Next  *Node
}

func loopDetection(ll LinkedList) *Node {
	curr := ll.Head
	addressMap := make(map[*Node]bool)
	for curr != nil {
		if _, ok := addressMap[curr]; ok {
			return curr
		}
		addressMap[curr] = true
		curr = curr.Next
	}
	return nil
}

func main() {
	loopStartNode := &Node{3, nil}

	input := LinkedList{&Node{1, &Node{2, loopStartNode}}}
	loopStartNode.Next = &Node{4, &Node{5, &Node{6, loopStartNode}}}
	fmt.Println(loopDetection(input))

	input2 := LinkedList{&Node{1, &Node{2, &Node{3, nil}}}}
	fmt.Println(loopDetection(input2))
}
