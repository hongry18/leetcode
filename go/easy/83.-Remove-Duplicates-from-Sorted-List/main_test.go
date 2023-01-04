package main

import "testing"

func TestXxx(t *testing.T) {
	n := &ListNode{}
	tmp := n

	tmp.Val = 1
	tmp.Next = &ListNode{}
	tmp = tmp.Next

	tmp.Val = 1
	tmp.Next = &ListNode{}
	tmp = tmp.Next

	tmp.Val = 2

	res := deleteDuplicates(n)
	for {
		t.Log(res.Val)
		res = res.Next
		if res == nil {
			break
		}
	}
}
