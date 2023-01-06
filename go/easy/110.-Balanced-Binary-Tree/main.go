package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	return getHeight(root) != -1
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left, right := 0, 0
	if left = getHeight(root.Left); left == -1 {
		return left
	}

	if right = getHeight(root.Right); right == -1 {
		return right
	}

	if abs(left-right) > 1 {
		return -1
	}
	return max(left, right) + 1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
