package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		n := &TreeNode{Val: 3}
		n.Left = &TreeNode{Val: 9}
		n.Right = &TreeNode{Val: 20}
		n.Right.Left = &TreeNode{Val: 15}
		n.Right.Right = &TreeNode{Val: 7}

		assert.Equal(t, isBalanced(n), true)
	})
}
