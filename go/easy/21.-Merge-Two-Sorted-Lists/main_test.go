package main

import (
	"testing"
)

func TestXxx(t *testing.T) {
	list1 := &ListNode{}
	list2 := &ListNode{}
	n := list1
	n.Val = 1
	n.Next = &ListNode{}

	n = n.Next
	n.Val = 2
	n.Next = &ListNode{}

	n = n.Next
	n.Val = 4

	n = list2
	n.Val = 1
	n.Next = &ListNode{}

	n = n.Next
	n.Val = 3
	n.Next = &ListNode{}

	n = n.Next
	n.Val = 4

	result := mergeTwoLists(list1, list2)

	for {
		t.Log(result.Val)
		if result.Next == nil {
			break
		}
		result = result.Next
	}
}
