package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	q1 := []*TreeNode{p}
	q2 := []*TreeNode{q}

	for len(q1) != 0 && len(q2) != 0 {
		n1 := q1[0]
		q1 = q1[1:]

		n2 := q2[0]
		q2 = q2[1:]

		if n1 == nil && n2 == nil {
			continue
		}

		if n1 == nil && n2 != nil || n1 != nil && n2 == nil {
			return false
		}

		if n1.Val != n2.Val {
			return false
		}
		q1 = append(q1, n1.Left, n1.Right)
		q2 = append(q2, n2.Left, n2.Right)
	}
	return len(q1) == 0 && len(q2) == 0
}

// func isSameTree(p *TreeNode, q *TreeNode) bool {
// 	if p == nil || q == nil {
// 		return p == q
// 	}

// 	if p.Val != q.Val {
// 		return false
// 	}

// 	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
// }
