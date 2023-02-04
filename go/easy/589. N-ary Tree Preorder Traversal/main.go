package main

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	if root == nil {
		return []int{}
	}
	var res = []int{root.Val}
	p(root, &res)
	return res
}

func p(parent *Node, res *[]int) {
	for _, n := range parent.Children {
		if n == nil {
			return
		}
		*res = append(*res, n.Val)
		p(n, res)
	}
}
