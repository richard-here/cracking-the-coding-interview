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

// Using buffer (hashmap)
func removeDuplicates(ll LinkedList) LinkedList {
	hashmap := make(map[int]bool)
	var prev *Node
	curr := ll.Head
	for curr != nil {
		if _, ok := hashmap[curr.Value]; ok {
			prev.Next = curr.Next
		} else {
			hashmap[curr.Value] = true
			prev = curr
		}
		curr = curr.Next
	}
	return ll
}

// Without extra buffer (space), nested loop is needed
// Worst case is O(n^2)
func removeDuplicatesWithoutBuffer(ll LinkedList) LinkedList {
	curr := ll.Head
	for curr != nil {
		valueToRemove := curr.Value
		prev := curr
		nodeToCheck := curr.Next
		for nodeToCheck != nil {
			if nodeToCheck.Value == valueToRemove {
				prev.Next = nodeToCheck.Next
			} else {
				prev = nodeToCheck
			}
			nodeToCheck = nodeToCheck.Next
		}
		curr = curr.Next
	}
	return ll
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func main() {
	input := LinkedList{Head: &Node{1, &Node{2, &Node{3, &Node{4, &Node{1, &Node{2, &Node{2, &Node{3, &Node{5, nil}}}}}}}}}}
	fmt.Println(prettyPrint(removeDuplicates(input)))
	fmt.Println(prettyPrint(removeDuplicatesWithoutBuffer(input)))
	input2 := LinkedList{Head: &Node{1, nil}}
	fmt.Println(prettyPrint(removeDuplicates(input2)))
	input3 := LinkedList{nil}
	fmt.Println(prettyPrint(removeDuplicatesWithoutBuffer(input3)))
}
