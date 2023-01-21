package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	var ar = []*TreeNode{root}

	var sum int
	var m = make(map[*TreeNode]bool)
	for len(ar) != 0 {
		n := ar[0]
		ar = ar[1:]

		if n.Left == nil && n.Right == nil {
			if _, ok := m[n]; ok {
				sum += n.Val
			}
		}

		if n.Left != nil {
			m[n.Left] = true
			ar = append(ar, n.Left)
		}

		if n.Right != nil {
			ar = append(ar, n.Right)
		}
	}

	return sum
}
