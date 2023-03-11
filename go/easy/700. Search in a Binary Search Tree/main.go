package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	var ar = []*TreeNode{root}
	for len(ar) != 0 {
		n := ar[0]
		ar = ar[1:]

		if n.Val == val {
			return n
		}

		if n.Val > val && n.Left != nil {
			ar = append(ar, n.Left)
		} else if n.Val < val && n.Right != nil {
			ar = append(ar, n.Right)
		}
	}
	return nil
}
