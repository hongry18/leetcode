package main

import "testing"

func TestXxx(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		r := &TreeNode{Val: 1}
		r.Left = &TreeNode{Val: 2}
		r.Right = &TreeNode{Val: 3}
		r.Left.Left = &TreeNode{Val: 4}
		r.Left.Right = &TreeNode{Val: 5}

		t.Log(depth(r.Left))
		t.Log(depth(r.Right))
	})
}
