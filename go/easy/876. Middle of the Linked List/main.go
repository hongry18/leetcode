package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	cnt := 0
	tmp := head
	for tmp != nil {
		cnt++
		tmp = tmp.Next
	}

	for i := 0; i < cnt/2; i++ {
		head = head.Next
	}

	return head
}

func optimizeMiddleNode(head *ListNode) *ListNode {
	fast, slow := head, head

	for fast != nil {
		fast = fast.Next
		if fast == nil {
			break
		}

		fast = fast.Next
		slow = slow.Next
	}

	return slow
}
