package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var sum int
	var a = []*TreeNode{root}
	for len(a) != 0 {
		n := a[0]
		a = a[1:]

		var l, r = sum(n.Left), sum(n.Right)
		sum = abs(l - r)
	}
	return 0
}

func sum(r *TreeNode) int {
	if r == nil {
		return 0
	}
	var s int
	var a = []*TreeNode{r}
	for len(a) != 0 {
		n := a[0]
		a = a[1:]
		s += n.Val
		if n.Left != nil {
			a = append(a, n.Left)
		}
		if n.Right != nil {
			a = append(a, n.Right)
		}
	}
	return s
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}
