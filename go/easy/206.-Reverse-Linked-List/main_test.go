package main

import "testing"

func TestXxx(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		n := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  5,
							Next: &ListNode{Val: 6},
						},
					},
				},
			},
		}

		r := reverseList(n)
		for r != nil {
			t.Log(r)
			r = r.Next
		}
	})
}
