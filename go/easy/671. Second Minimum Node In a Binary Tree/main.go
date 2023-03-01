package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findSecondMinimumValue(root *TreeNode) int {
	var m = make(map[int]bool)
	var ar = []*TreeNode{root}

	for len(ar) != 0 {
		n := ar[0]
		ar = ar[1:]
		if n == nil {
			continue
		}
		m[n.Val] = true
		ar = append(ar, n.Left, n.Right)
	}

	if len(m) < 2 {
		return -1
	}

	var m1, m2 int = math.MaxInt, math.MaxInt
	for k := range m {
		if k < m1 {
			m2 = m1
			m1 = k
			continue
		}

		if k < m2 {
			m2 = k
		}
	}
	return m2
}
