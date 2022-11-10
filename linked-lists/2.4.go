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

func (ll *LinkedList) appendToLl(n *Node) {
	if ll.Head == nil {
		ll.Head = n
		return
	}
	curr := ll.Head
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = n
}

func partition(ll LinkedList, n int) LinkedList {
	// Left and right partition
	lp := LinkedList{nil}
	rp := LinkedList{nil}
	var m *Node
	var lastLpNode *Node
	var prev *Node
	curr := ll.Head

	// Place nodes on lp or rp, depending on the value of the nodes
	for curr != nil {
		currValue := curr.Value
		prev, curr = curr, curr.Next
		prev.Next = nil
		if currValue < n {
			lp.appendToLl(prev)
			lastLpNode = prev
		} else if m == nil && n == prev.Value {
			m = prev
		} else {
			rp.appendToLl(prev)
		}
	}

	// If no node is on the left partition, make the middle value as the head
	// Otherwise, the middle value will be the bridge between the left and right partitions
	if lastLpNode == nil {
		ll.Head = m
		m.Next = rp.Head
	} else {
		ll.Head = lp.Head
		lastLpNode.Next = m
		m.Next = rp.Head
	}

	return ll
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", " ")
	return string(s)
}

func main() {
	input := LinkedList{&Node{3, &Node{5, &Node{8, &Node{5, &Node{10, &Node{2, &Node{1, nil}}}}}}}}
	num := 5
	fmt.Print(prettyPrint(partition(input, num)))
}
