package main

import "testing"

func TestXxx(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		n := &TreeNode{Val: -9}
		n.Left = &TreeNode{Val: -3}
		n.Left.Right = &TreeNode{Val: 4}
		n.Left.Left = &TreeNode{Val: -6}
		n.Right = &TreeNode{Val: 2}
		n.Right.Left = &TreeNode{Val: 4}
		n.Right.Right = &TreeNode{Val: 0}
		n.Right.Left.Left = &TreeNode{Val: -5}

		sumOfLeftLeaves(n)
	})
}
