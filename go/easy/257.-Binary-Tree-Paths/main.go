package main

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type NodePath struct {
	Node *TreeNode
	Path []int
}

func binaryTreePaths(root *TreeNode) []string {
	var res []string
	ar := []NodePath{{Node: root}}

	var sb strings.Builder
	for len(ar) != 0 {
		n := ar[len(ar)-1]
		ar = ar[:len(ar)-1]

		if n.Node.Left == nil && n.Node.Right == nil {
			for _, v := range n.Path {
				sb.WriteString(strconv.Itoa(v))
				sb.WriteString("->")
			}
			sb.WriteString(strconv.Itoa(n.Node.Val))
			res = append(res, sb.String())
			sb.Reset()
		}

		if n.Node.Right != nil {
			ar = append(ar, NodePath{Node: n.Node.Right, Path: append(n.Path, n.Node.Val)})
		}

		if n.Node.Left != nil {
			ar = append(ar, NodePath{Node: n.Node.Left, Path: append(n.Path, n.Node.Val)})
		}
	}

	return res
}
