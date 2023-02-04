package main

import (
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	var ar = []*TreeNode{root}
	var ar2 []int

	for len(ar) != 0 {
		n := ar[0]
		ar = ar[1:]
		ar2 = append(ar2, n.Val)
		if n.Left != nil {
			ar = append(ar, n.Left)
		}

		if n.Right != nil {
			ar = append(ar, n.Right)
		}
	}
	sort.Ints(ar2)
	var res int = 1000000
	for i := 0; i < len(ar2)-1; i++ {
		t := ar2[i+1] - ar2[i]
		if t < res {
			res = t
		}
		if t == 1 {
			break
		}
	}
	return res
}
