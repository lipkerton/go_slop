package main

import "fmt"

var numbers_1 []int = []int{1, 2, 3, 4}
var numbers_2 []int = []int{1, 2, 3, 4}

type ListNode struct {
	Val  int
	Next *ListNode
}

func makeLinkedList(index int, n []int) *ListNode {
	if index >= len(n) {
		return nil
	}
	var result *ListNode = &ListNode{Val: n[index], Next: makeLinkedList(index+1, n)}
	return result
}

var ln_1 *ListNode = &ListNode{Val: numbers_1[0], Next: makeLinkedList(1, numbers_1)}
var ln_2 *ListNode = &ListNode{Val: numbers_2[0], Next: makeLinkedList(1, numbers_2)}

func parseLinkedList(n_1 *ListNode, n_2 *ListNode) *ListNode {
	if n_1 == nil && n_2 == nil {
		return nil
	}
	if n_1 == nil {
		n_1 = &ListNode{Val: 0, Next: nil}
	}
	if n_2 == nil {
		n_2 = &ListNode{Val: 0, Next: nil}
	}
	var result *ListNode
	if n_1.Val == n_2.Val {
		result = *n_1
		result.Next = n_2
		result.Next.Next = parseLinkedList(n_1.Next, n_2.Next)
		return result.Next.Next
	} else if n_1.Val > n_2.Val {
		result = n_2
		result.Next = parseLinkedList(n_1, n_2.Next)
		return result.Next
	} else if n_2.Val > n_1.Val {
		result = n_1
		result.Next = parseLinkedList(n_2, n_1.Next)
		return result.Next
	}
	return result
}

func main() {
	fmt.Println(ln_1.Val, ln_2.Val, ln_1.Next.Val, ln_2.Next.Val)
	var result *ListNode = &ListNode{Val: 0, Next: parseLinkedList(ln_1, ln_2)}
	fmt.Println(result)
}
