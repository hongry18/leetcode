package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(root *TreeNode) string {
	var q = []*TreeNode{root}

	for len(q) != 0 {
		n := q[len(q)-1]
		q = q[:len(q)-1]

		if n.Right != nil {
			q = append(q, n.Right)
		}

		if n.Left != nil {
			q = append(q, n.Left)
		}
	}
	return ""
}
