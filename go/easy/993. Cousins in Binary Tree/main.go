package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x, y int) bool {
	var q = []*TreeNode{root}

	for len(q) > 0 {
		var res int
		var t []*TreeNode
		for i := 0; i < len(q); i++ {
			if q[i] == nil {
				continue
			}

			if q[i].Left != nil && (q[i].Left.Val == x || q[i].Left.Val == y) {
				res++
			} else if q[i].Right != nil && (q[i].Right.Val == x || q[i].Right.Val == y) {
				res++
			}

			t = append(t, q[i].Left, q[i].Right)
		}

		if res == 2 {
			return true
		}
		q = t
	}
	return false
}
