package main

import "testing"

func TestXxx(t *testing.T) {
	r := &TreeNode{Val: 0}
	r.Right = &TreeNode{Val: 2}
	r.Right.Left = &TreeNode{Val: 1}
	r.Right.Right = &TreeNode{Val: 2}
	r.Right.Right.Right = &TreeNode{Val: 3}
	r.Right.Right.Right.Right = &TreeNode{Val: 3}
	findMode(r)
}
