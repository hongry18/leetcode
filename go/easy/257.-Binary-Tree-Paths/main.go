package main

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	ar := []*TreeNode{root}

	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	}

	for len(ar) != 0 {
		n := ar[len(ar)-1]
		ar = ar[:len(ar)-1]
		if n == nil {
			continue
		}

		if n.Left == nil && n.Right == nil {
			// 처리
		}

		if n.Left != nil {
			ar = append(ar, n.Left)
		}

		if n.Right != nil {
			ar = append(ar, n.Right)
		}
	}

	return []string{}
}
