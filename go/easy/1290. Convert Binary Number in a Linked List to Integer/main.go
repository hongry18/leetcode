package main

import "fmt"

func main() {
	head := gen([]int{1, 0, 1, 0, 0, 1, 1, 0, 1, 1, 1, 0, 1, 0, 0, 1, 1, 0, 1, 1, 1, 0, 1, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1})
	fmt.Println(getDecimalValue(head))
}

func gen(ar []int) *ListNode {
	head := &ListNode{Val: ar[0]}
	n := head
	for i := 1; i < len(ar); i++ {
		n.Next = &ListNode{Val: ar[i]}
		n = n.Next
	}
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	var res int
	for head != nil {
		res = (res << 1) | head.Val
		head = head.Next
	}
	return res
}
