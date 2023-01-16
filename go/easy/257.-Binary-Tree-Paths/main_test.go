package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		r := &TreeNode{Val: 1}
		r.Left = &TreeNode{Val: 2}
		r.Left.Right = &TreeNode{Val: 5}
		r.Right = &TreeNode{Val: 3}

		assert.Equal(t, binaryTreePaths(r), []string{"1->2->5", "1->3"})
	})
}
