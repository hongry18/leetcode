package main

import "fmt"

func main() {
	r := &TreeNode{Val: 1}
	r.Left = &TreeNode{Val: 2}
	r.Left.Left = &TreeNode{Val: 4}
	r.Right = &TreeNode{Val: 3}
	fmt.Println(tree2str(r))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(root *TreeNode) string {
	if root == nil {
		return ""
	}

	var l, r string

	if root.Left != nil {
		l = fmt.Sprintf("(%s)", tree2str(root.Left))
	}

	if root.Right != nil {
		l = fmt.Sprintf("(%s)", tree2str(root.Left))
		r = fmt.Sprintf("(%s)", tree2str(root.Right))
	}

	return fmt.Sprintf("%d%s%s", root.Val, l, r)
}

func tree2str2(root *TreeNode) string {
	if root.Left == nil && root.Right == nil {
		return fmt.Sprintf("%d", root.Val)
	}
	return fmt.Sprintf("%d%s%s", root.Val, dfs(root.Left, true), dfs(root.Right, false))
}

func dfs(n *TreeNode, isLeft bool) string {
	if n == nil {
		if isLeft {
			return "()"
		}
		return ""
	}

	if n.Left == nil && n.Right == nil {
		return fmt.Sprintf("(%d)", n.Val)
	}

	return fmt.Sprintf("(%d%s%s)", n.Val, dfs(n.Left, true), dfs(n.Right, false))
}
