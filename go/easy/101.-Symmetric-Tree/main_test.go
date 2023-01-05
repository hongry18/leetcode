package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		n1 := &TreeNode{Val: 1}
		n1.Left = &TreeNode{Val: 2}
		n1.Left.Left = &TreeNode{Val: 3}
		n1.Left.Right = &TreeNode{Val: 4}
		n1.Right = &TreeNode{Val: 2}
		n1.Right.Left = &TreeNode{Val: 4}
		n1.Right.Right = &TreeNode{Val: 3}

		assert.Equal(t, isSymmetric(n1), true)
	})

	t.Run("test2", func(t *testing.T) {
		n1 := &TreeNode{Val: 1}
		n1.Left = &TreeNode{Val: 2}
		n1.Left.Right = &TreeNode{Val: 3}
		n1.Right = &TreeNode{Val: 2}
		n1.Right.Right = &TreeNode{Val: 3}

		assert.Equal(t, isSymmetric(n1), false)
	})

	t.Run("test3", func(t *testing.T) {
		n1 := &TreeNode{Val: 9}

		n1.Left = &TreeNode{Val: -42}
		n1.Left.Right = &TreeNode{Val: 76}
		n1.Left.Right.Right = &TreeNode{Val: 13}

		n1.Right = &TreeNode{Val: -42}
		n1.Right.Left = &TreeNode{Val: 76}
		n1.Right.Left.Right = &TreeNode{Val: 13}

		assert.Equal(t, isSymmetric(n1), false)
	})
}
