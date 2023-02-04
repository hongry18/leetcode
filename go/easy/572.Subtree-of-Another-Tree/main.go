package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	return dfs(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func dfs(r, s *TreeNode) bool {
	if r == nil && s == nil {
		return true
	}

	if r == nil || s == nil {
		return false
	}
	if r.Val != s.Val {
		return false
	}
	return dfs(r.Left, s.Left) && dfs(r.Right, s.Right)
}
