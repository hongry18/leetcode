package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	return depth(root.Left) + depth(root.Right)
}

func depth(r *TreeNode) int {
	if r == nil {
		return 0
	}
	var cnt int
	var ar = []*TreeNode{r}
	for len(ar) != 0 {
		cnt++
		var ar2 []*TreeNode
		for _, n := range ar {
			if n.Left != nil {
				ar2 = append(ar2, n.Left)
			}
			if n.Right != nil {
				ar2 = append(ar2, n.Right)
			}
		}
		ar = ar2
	}

	return cnt
}
