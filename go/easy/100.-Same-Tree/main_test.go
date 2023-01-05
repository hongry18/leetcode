package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		n1 := &TreeNode{Val: 1}
		n1.Left = &TreeNode{Val: 2}
		n1.Right = &TreeNode{Val: 3}

		n2 := &TreeNode{Val: 1}
		n2.Left = &TreeNode{Val: 2}
		n2.Right = &TreeNode{Val: 3}

		assert.Equal(t, isSameTree(n1, n2), true)
	})

	t.Run("test2", func(t *testing.T) {
		n1 := &TreeNode{Val: 1}
		n1.Left = &TreeNode{Val: 2}

		n2 := &TreeNode{Val: 1}
		n2.Right = &TreeNode{Val: 2}

		assert.Equal(t, isSameTree(n1, n2), false)
	})

	t.Run("test3", func(t *testing.T) {
		n1 := &TreeNode{Val: 1}
		n1.Right = &TreeNode{Val: 2}

		n2 := &TreeNode{Val: 1}
		n2.Left = &TreeNode{Val: 2}

		assert.Equal(t, isSameTree(n1, n2), false)
	})

	t.Run("test4", func(t *testing.T) {
		n1 := &TreeNode{Val: 1}
		n1.Left = &TreeNode{Val: 2}
		n1.Right = &TreeNode{Val: 1}

		n2 := &TreeNode{Val: 1}
		n2.Left = &TreeNode{Val: 1}
		n2.Right = &TreeNode{Val: 2}

		assert.Equal(t, isSameTree(n1, n2), false)
	})

	t.Run("test4", func(t *testing.T) {
		n1 := &TreeNode{Val: 1}

		n2 := &TreeNode{Val: 1}
		n2.Right = &TreeNode{Val: 2}

		assert.Equal(t, isSameTree(n1, n2), false)
	})
}
