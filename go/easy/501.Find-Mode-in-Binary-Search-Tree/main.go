package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	var res []int
	var ar = []*TreeNode{root}
	var max int
	var m = map[int]int{}

	for len(ar) != 0 {
		n := ar[len(ar)-1]
		ar = ar[:len(ar)-1]
		m[n.Val]++
		if max < m[n.Val] {
			max = m[n.Val]
		}

		if n.Left != nil {
			ar = append(ar, n.Left)
		}

		if n.Right != nil {
			ar = append(ar, n.Right)
		}
	}

	for i, v := range m {
		if v == max {
			res = append(res, i)
		}
	}
	return res
}
