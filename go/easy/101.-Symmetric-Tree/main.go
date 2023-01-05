package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}

	// recursive
	// return recursive(root.Left, root.Right)

	ar := []*TreeNode{root.Left, root.Right}

	for len(ar) != 0 {
		l := ar[0]
		r := ar[1]

		ar = ar[2:]
		if l == nil && r == nil {
			continue
		}

		if (l == nil && r != nil) || (l != nil && r == nil) {
			return false
		}

		if l.Val != r.Val {
			return false
		}

		ar = append(ar, l.Left, r.Right, l.Right, r.Left)
	}

	return true
}

func recursive(l *TreeNode, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}

	if l == nil || r == nil || l.Val != r.Val {
		return false
	}

	return recursive(l.Left, r.Right) && recursive(l.Right, r.Left)
}
