package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	n := &TreeNode{Val: 0}
	n.Left = &TreeNode{Val: 2}
	n.Left.Left = &TreeNode{Val: 1}
	n.Left.Left.Left = &TreeNode{Val: 5}
	n.Left.Left.Right = &TreeNode{Val: 1}

	n.Right = &TreeNode{Val: 4}
	n.Right.Left = &TreeNode{Val: 3}
	n.Right.Left.Right = &TreeNode{Val: 6}
	n.Right.Right = &TreeNode{Val: -1}
	n.Right.Right.Right = &TreeNode{Val: 8}

	assert.Equal(t, maxDepth(n), 4)
}

func Test2(t *testing.T) {
	n := &TreeNode{Val: 0}
	assert.Equal(t, maxDepth(n), 1)
}
