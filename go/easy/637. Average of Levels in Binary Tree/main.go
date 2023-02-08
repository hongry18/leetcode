package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	r := &TreeNode{Val: 3}
	r.Left = &TreeNode{Val: 9}
	r.Right = &TreeNode{Val: 20}
	r.Right.Left = &TreeNode{Val: 15}
	r.Right.Right = &TreeNode{Val: 7}
	fmt.Println(averageOfLevels(r))
}

func averageOfLevels(root *TreeNode) []float64 {
	var a = []*TreeNode{root}
	var res []float64
	for len(a) > 0 {
		var b []*TreeNode
		var cnt, sum float64
		for _, n := range a {
			cnt++
			sum += float64(n.Val)
			if n.Left != nil {
				b = append(b, n.Left)
			}
			if n.Right != nil {
				b = append(b, n.Right)
			}
		}
		res = append(res, sum/cnt)
		a = b
	}
	return res
}
