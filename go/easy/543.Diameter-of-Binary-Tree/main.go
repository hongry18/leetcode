package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	n := 0
	dfs(root, &n)
	return n
}

func dfs(root *TreeNode, n *int) int {
	if root == nil {
		return 0
	}
	l := dfs(root.Left, n)
	r := dfs(root.Right, n)
	if l+r > *n {
		*n = l + r
	}
	return 1 + max(l, r)

}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
