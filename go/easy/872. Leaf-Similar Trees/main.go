package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var list = []*TreeNode{root1}
	var ar1, ar2 []int

	for len(list) != 0 {
		n := list[len(list)-1]
		list = list[:len(list)-1]

		if n.Left == nil && n.Right == nil {
			ar1 = append(ar1, n.Val)
		}

		if n.Right != nil {
			list = append(list, n.Right)
		}
		if n.Left != nil {
			list = append(list, n.Left)
		}
	}

	list = []*TreeNode{root2}
	for len(list) != 0 {
		n := list[len(list)-1]
		list = list[:len(list)-1]

		if n.Left == nil && n.Right == nil {
			ar2 = append(ar2, n.Val)
		}

		if n.Right != nil {
			list = append(list, n.Right)
		}
		if n.Left != nil {
			list = append(list, n.Left)
		}
	}

	if len(ar1) != len(ar2) {
		return false
	}

	for i := range ar1 {
		if ar1[i] != ar2[i] {
			return false
		}
	}

	return true
}
