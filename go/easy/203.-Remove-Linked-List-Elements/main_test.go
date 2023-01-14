package main

import (
	"testing"
)

func TestXxx(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		head := &ListNode{Val: 1}
		n := head
		head.Next = &ListNode{Val: 2}

		n = n.Next
		n.Next = &ListNode{Val: 6}
		n = n.Next
		n.Next = &ListNode{Val: 3}
		n = n.Next
		n.Next = &ListNode{Val: 4}
		n = n.Next
		n.Next = &ListNode{Val: 5}
		n = n.Next
		n.Next = &ListNode{Val: 6}

		result := removeElements(head, 6)

		for result != nil {
			t.Log(result.Val)
			result = result.Next
		}
	})
}
