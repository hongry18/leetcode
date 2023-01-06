package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 자식이 없는 노드가 있을때 반환
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left, right := minDepth(root.Left), minDepth(root.Right)
	if left == 0 || right == 0 {
		return left + right + 1
	}
	if left < right {
		return left + 1
	}
	return right + 1

	// q := []*TreeNode{root}
	// depth := 0

	// for len(q) != 0 {
	// 	depth++

	// 	for _, n := range q {
	// 		q = q[1:]
	// 		if n.Left == nil && n.Right == nil {
	// 			return depth
	// 		}
	// 		if n.Left != nil {
	// 			q = append(q, n.Left)
	// 		}
	// 		if n.Right != nil {
	// 			q = append(q, n.Right)
	// 		}
	// 	}
	// }

	// return depth
}
