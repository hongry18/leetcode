package main

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	if root == nil {
		return []int{}
	}

	var ar = []*Node{root}
	var res []int
	for len(ar) != 0 {
		n := ar[len(ar)-1]
		ar = ar[:len(ar)-1]
		res = append(res, n.Val)
		for i := len(n.Children) - 1; i >= 0; i-- {
			ar = append(ar, n.Children[i])
		}
	}
	return res
}
