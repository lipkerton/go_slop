package main
import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func recurList(number int, cur *ListNode) *ListNode {
	if number == 10 {
		result := ListNode{Val: number, Next: nil}
		return &result
	}
	cur.Next = &ListNode{Val: number, Next: nil}
	return recurList(number + 1, cur.Next)
}

func main() {
	head := ListNode{Val: 1, Next: nil}
	recurList(head.Val + 1, &head)
	fmt.Println(head)
	fmt.Println(head.Next.Val)
}