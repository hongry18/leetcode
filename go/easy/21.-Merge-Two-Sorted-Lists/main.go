package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	result := &ListNode{}

	t := result

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			t.Val = list1.Val
			list1 = list1.Next
		} else {
			t.Val = list2.Val
			list2 = list2.Next
		}
		t.Next = &ListNode{}
		t = t.Next
	}

	for list1 != nil {
		t.Val = list1.Val
		list1 = list1.Next
		if list1 == nil {
			break
		}
		t.Next = &ListNode{}
		t = t.Next
	}

	for list2 != nil {
		t.Val = list2.Val
		list2 = list2.Next
		if list2 == nil {
			break
		}
		t.Next = &ListNode{}
		t = t.Next
	}

	return result
}
