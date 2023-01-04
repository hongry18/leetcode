package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	n := &TreeNode{Val: 1}
	n.Right = &TreeNode{Val: 2}
	n.Right.Left = &TreeNode{Val: 3}

	assert.Equal(t, inorderTraversal(n), []int{1, 3, 2})
}
