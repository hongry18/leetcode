package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
}

func sumRootToLeaf(root *TreeNode) int {
	l := []*TreeNode{root}
	var sum int
	for len(l) > 0 {
		n := l[0]
		l = l[1:]

		if n.Left == nil && n.Right == nil {
			sum += n.Val
			continue
		}

		if n.Left != nil {
			n.Left.Val = n.Val*2 + n.Left.Val
			l = append(l, n.Left)
		}

		if n.Right != nil {
			n.Right.Val = n.Val*2 + n.Right.Val
			l = append(l, n.Right)
		}
	}
	return sum
}
