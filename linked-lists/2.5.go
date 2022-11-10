package main

import (
	"encoding/json"
	"fmt"
	"math"
)

type LinkedList struct {
	Head *Node
}

type Node struct {
	Value int
	Next  *Node
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", " ")
	return string(s)
}

func getReverseOrderNumber(ll LinkedList) int {
	number := 0
	curr := ll.Head
	pow := 0
	for curr != nil {
		number += curr.Value * int(math.Pow(10, float64(pow)))
		pow++
		curr = curr.Next
	}
	return number
}

func getForwardOrderNumber(ll LinkedList) int {
	number := 0
	curr := ll.Head
	pow := 0
	for curr != nil {
		number *= 10
		number += curr.Value
		pow++
		curr = curr.Next
	}
	return number
}

func sumListReverseOrder(ll1 LinkedList, ll2 LinkedList) int {
	return getReverseOrderNumber(ll1) + getReverseOrderNumber(ll2)
}

func sumListForwardOrder(ll1 LinkedList, ll2 LinkedList) int {
	return getForwardOrderNumber(ll1) + getForwardOrderNumber(ll2)
}

func main() {
	ll1 := LinkedList{&Node{7, &Node{1, &Node{6, nil}}}}
	ll2 := LinkedList{&Node{5, &Node{9, &Node{2, nil}}}}
	fmt.Println(sumListReverseOrder(ll1, ll2))

	ll3 := LinkedList{&Node{6, &Node{1, &Node{7, nil}}}}
	ll4 := LinkedList{&Node{2, &Node{9, &Node{5, nil}}}}
	fmt.Println(sumListForwardOrder(ll3, ll4))
}
