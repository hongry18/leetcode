package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	r := &TreeNode{Val: 5}
	r.Left = &TreeNode{Val: 3}
	r.Left.Left = &TreeNode{Val: 2}
	r.Left.Left.Left = &TreeNode{Val: 1}
	r.Left.Right = &TreeNode{Val: 4}
	r.Right = &TreeNode{Val: 6}
	r.Right.Right = &TreeNode{Val: 8}
	r.Right.Right.Left = &TreeNode{Val: 7}
	r.Right.Right.Right = &TreeNode{Val: 9}

	i := increasingBST(r)

	for i != nil {
		fmt.Println(i.Val)
		i = i.Right
	}
}

func increasingBST(root *TreeNode) *TreeNode {
	ar := []int{}
	bst(root, &ar)
	r := &TreeNode{Val: ar[0]}
	t := r
	for i := 1; i < len(ar); i++ {
		c := &TreeNode{Val: ar[i]}
		t.Right = c
		t = t.Right
	}
	return r
}

func bst(r *TreeNode, ar *[]int) {
	if r == nil {
		return
	}
	bst(r.Left, ar)
	*ar = append(*ar, r.Val)
	bst(r.Right, ar)
}
