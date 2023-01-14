package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	head.Next = removeElements(head.Next, val)

	if head.Val == val {
		return head.Next
	}
	return head
}

func removeElements1(head *ListNode, val int) *ListNode {
	result := &ListNode{}
	t := result

	for head != nil {
		if head.Val != val {
			t.Next = &ListNode{Val: head.Val}
			t = t.Next
		}
		head = head.Next
	}

	return result.Next
}
