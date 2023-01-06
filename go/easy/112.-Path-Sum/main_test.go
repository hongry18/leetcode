package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		n := &TreeNode{Val: 5}
		n.Left = &TreeNode{Val: 4}
		n.Left.Left = &TreeNode{Val: 11}
		n.Left.Left.Left = &TreeNode{Val: 7}
		n.Left.Left.Right = &TreeNode{Val: 2}
		n.Right = &TreeNode{Val: 8}
		n.Right.Left = &TreeNode{Val: 13}
		n.Right.Right = &TreeNode{Val: 4}
		n.Right.Right.Right = &TreeNode{Val: 1}
		assert.Equal(t, hasPathSum(n, 22), true)
	})
}
