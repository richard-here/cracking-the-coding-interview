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

func intersectingNode(ll1 LinkedList, ll2 LinkedList) *Node {
	curr := ll1.Head
	addressMap := make(map[*Node]bool)
	for curr != nil {
		addressMap[curr] = true
		curr = curr.Next
	}
	curr = ll2.Head
	for curr != nil {
		if _, ok := addressMap[curr]; ok {
			return curr
		}
		curr = curr.Next
	}
	return nil
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", " ")
	return string(s)
}

func main() {
	node := &Node{3, &Node{4, &Node{3, &Node{5, nil}}}}
	ll1 := LinkedList{&Node{1, &Node{2, &Node{3, node}}}}
	ll2 := LinkedList{&Node{2, node}}
	fmt.Println(prettyPrint(intersectingNode(ll1, ll2)))

	ll3 := LinkedList{&Node{1, &Node{2, nil}}}
	ll4 := LinkedList{&Node{2, nil}}
	fmt.Println(prettyPrint(intersectingNode(ll3, ll4)))
}
