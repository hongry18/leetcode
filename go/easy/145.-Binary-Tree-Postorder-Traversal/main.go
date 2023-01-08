package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	// travel(root, &res)
	// return res

	ar := []*TreeNode{}
	cur := root

	for cur != nil || len(ar) != 0 {
		if cur == nil {
			n := ar[len(ar)-1]
			ar = ar[:len(ar)-1]
			cur = n.Left
		} else {
			ar = append(ar, cur)
			res = append([]int{cur.Val}, res...)
			cur = cur.Right
		}
	}

	return res
}

func travel(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}

	travel(root.Left, ans)
	travel(root.Right, ans)
	*ans = append(*ans, root.Val)
}
