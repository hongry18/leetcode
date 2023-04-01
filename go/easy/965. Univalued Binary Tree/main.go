package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	ar := []*TreeNode{root}

	for len(ar) > 0 {
		n := ar[0]
		ar = ar[1:]

		if n.Left != nil {
			if n.Val != n.Left.Val {
				return false
			}
			ar = append(ar, n.Left)
		}
		if n.Right != nil {
			if n.Val != n.Right.Val {
				return false
			}
			ar = append(ar, n.Right)
		}
	}
	return true
}
