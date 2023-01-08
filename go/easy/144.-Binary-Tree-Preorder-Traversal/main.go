package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	travel(root, &res)
	return res

	ar := []*TreeNode{root}

	for len(ar) != 0 {
		n := ar[len(ar)-1]
		ar = ar[:len(ar)-1]

		if n == nil {
			continue
		}

		res = append(res, n.Val)
		ar = append(ar, n.Right, n.Left)
	}

	return res
}

func travel(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}

	*ans = append(*ans, root.Val)
	travel(root.Left, ans)
	travel(root.Right, ans)
	fmt.Println(root.Val)
}
