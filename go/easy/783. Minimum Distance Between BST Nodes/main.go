package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDiffInBST(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var min int = 1000000
	var ar = []*TreeNode{root}
	for len(ar) > 0 {
		n := ar[0]
		ar = ar[1:]

		t := cmp(n)
		if t < min {
			min = t
		}

		if n.Left != nil {
			ar = append(ar, n.Left)
		}

		if n.Right != nil {
			ar = append(ar, n.Right)
		}
	}

	return min
}

func cmp(r *TreeNode) int {
	if r == nil {
		return 0
	}

	var min int = 1000000
	var ar = []*TreeNode{r.Left, r.Right}
	for len(ar) > 0 {
		n := ar[0]
		ar = ar[1:]
		if n == nil {
			continue
		}

		t := abs(r.Val - n.Val)
		if t < min {
			min = t
		}
	}
	return min
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}
