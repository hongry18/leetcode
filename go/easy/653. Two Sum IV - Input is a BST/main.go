package main

import "fmt"

func main() {
	r := &TreeNode{Val: 5}
	r.Left = &TreeNode{Val: 3}
	r.Left.Left = &TreeNode{Val: 2}
	r.Left.Right = &TreeNode{Val: 4}
	r.Right = &TreeNode{Val: 6}
	r.Right.Right = &TreeNode{Val: 7}

	fmt.Println(findTarget(r, 28))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	var m = map[int]bool{}
	var ar = []*TreeNode{root}

	for len(ar) != 0 {
		n := ar[0]
		ar = ar[1:]

		if m[k-n.Val] {
			return true
		}
		m[n.Val] = true

		if n.Right != nil {
			ar = append(ar, n.Right)
		}
		if n.Left != nil {
			ar = append(ar, n.Left)
		}
	}
	return false
}
