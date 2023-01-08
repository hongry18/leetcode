package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		n := &TreeNode{Val: 1}
		n.Right = &TreeNode{Val: 2}
		n.Right.Left = &TreeNode{Val: 3}

		assert.Equal(t, preorderTraversal(n), []int{1, 2, 3})
	})

	t.Run("test2", func(t *testing.T) {
		n := &TreeNode{Val: 1}
		n.Left = &TreeNode{Val: 4}
		n.Right = &TreeNode{Val: 3}
		n.Left.Left = &TreeNode{Val: 2}

		assert.Equal(t, preorderTraversal(n), []int{1, 4, 2, 3})
	})

	t.Run("test3", func(t *testing.T) {
		n := &TreeNode{Val: 2}
		n.Left = &TreeNode{Val: 1}
		n.Right = &TreeNode{Val: 3}
		n.Left.Right = &TreeNode{Val: 4}

		assert.Equal(t, preorderTraversal(n), []int{2, 1, 4, 3})
	})
}
