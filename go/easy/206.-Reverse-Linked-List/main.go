package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var result *ListNode
	if head == nil {
		return result
	}

	for head != nil {
		n := &ListNode{Val: head.Val, Next: result}
		result = n
		head = head.Next
	}

	return result
}
