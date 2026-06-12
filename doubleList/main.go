package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func checkRecurs(cur *ListNode, numbers []int) *ListNode {
	if cur == nil {
		return nil
	}
	// проверка
	for i := 0; i < len(numbers); i++ {
		if cur.Val == numbers[i] {
			return checkRecurs(cur.Next, numbers)
		}
	}
	numbers = append(numbers, cur.Val)
	cur.Next = checkRecurs(cur.Next, numbers)
	return cur
}

func makeDoubleList(index int, numbers []int) *ListNode {
	if index == len(numbers)-1 {
		return &ListNode{Val: numbers[len(numbers)-1], Next: nil}
	}
	cur := &ListNode{
		Val:  numbers[index],
		Next: makeDoubleList(index+1, numbers),
	}
	return cur
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 5, 5, 4, 10}
	head := &ListNode{
		Val:  numbers[0],
		Next: makeDoubleList(1, numbers),
	}
	head.Next = checkRecurs(head.Next, []int{head.Val})
	cur := head
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Next
	}
}
