package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	a := 0
	dfs(root, &a)
	return a
}

func dfs(n *TreeNode, a *int) int {
	if n == nil {
		return 0
	}
	l := dfs(n.Left, a)
	r := dfs(n.Right, a)
	*a += abs(l - r)
	return l + r + n.Val
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}
