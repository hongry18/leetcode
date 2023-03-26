package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	var sum int
	var list = []*TreeNode{root}

	for len(list) > 0 {
		n := list[0]
		list = list[1:]

		if n.Val >= low && n.Val <= high {
			sum += n.Val
			if n.Left != nil {
				list = append(list, n.Left)
			}

			if n.Right != nil {
				list = append(list, n.Right)
			}
			continue
		} else if n.Val < low && n.Right != nil {
			list = append(list, n.Right)
		} else if n.Val > high && n.Left != nil {
			list = append(list, n.Left)
		}
	}

	return sum
}
