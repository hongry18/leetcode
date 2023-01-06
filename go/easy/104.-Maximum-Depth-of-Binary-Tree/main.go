package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	// if root == nil {
	// 	return 0
	// }

	// return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
	if root == nil {
		return 0
	}
	q := []*TreeNode{root}
	depth := 0

	for len(q) != 0 {
		depth++
		for _, n := range q {
			q = q[1:]
			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
		}
	}
	return depth
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
