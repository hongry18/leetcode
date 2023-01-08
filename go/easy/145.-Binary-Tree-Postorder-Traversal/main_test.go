package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	t.Run("test3", func(t *testing.T) {
		n := &TreeNode{Val: 8}
		n.Left = &TreeNode{Val: 4}
		n.Left.Left = &TreeNode{Val: 1}
		n.Left.Right = &TreeNode{Val: 3}
		n.Left.Right.Left = &TreeNode{Val: 2}
		n.Right = &TreeNode{Val: 7}
		n.Right.Left = &TreeNode{Val: 5}
		n.Right.Right = &TreeNode{Val: 6}

		assert.Equal(t, postorderTraversal(n), []int{1, 2, 3, 4, 5, 6, 7, 8})
	})
}
