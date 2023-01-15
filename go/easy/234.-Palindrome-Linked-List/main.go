package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	var ar []int

	for head != nil {
		ar = append(ar, head.Val)
		head = head.Next
	}

	for i := 0; i < len(ar)/2; i++ {
		if ar[i] != ar[len(ar)-1-i] {
			return false
		}
	}
	return true
}
