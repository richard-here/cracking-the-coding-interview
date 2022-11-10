package main

import "fmt"

type LinkedList struct {
	Head *Node
}

type Node struct {
	Value int
	Next  *Node
}

func palindrome(ll LinkedList) bool {
	if ll.Head == nil || ll.Head.Next == nil {
		return true
	}
	llSize := 0
	arr := []int{}
	curr := ll.Head
	for curr != nil {
		if len(arr) == 0 || curr.Value != arr[len(arr)-1] {
			arr = append(arr, curr.Value)
		} else {
			arr = arr[:len(arr)-1]
		}
		llSize++
		curr = curr.Next
	}

	if llSize%2 == 0 {
		return len(arr) == 0
	} else {
		isPalindrome := true
		left := arr[:len(arr)/2]
		right := arr[len(arr)/2+1:]
		for i := range left {
			if left[i] != right[len(right)-1-i] {
				isPalindrome = false
				break
			}
		}
		return isPalindrome
	}
}

func main() {
	ll1 := LinkedList{&Node{1, &Node{2, &Node{3, &Node{2, &Node{1, nil}}}}}}
	ll2 := LinkedList{&Node{1, &Node{2, &Node{2, &Node{1, nil}}}}}
	ll3 := LinkedList{&Node{1, &Node{2, &Node{3, &Node{4, nil}}}}}
	fmt.Println(palindrome(ll1))
	fmt.Println(palindrome(ll2))
	fmt.Println(palindrome(ll3))
}
