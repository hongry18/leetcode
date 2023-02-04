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
		n := ar[0]
		res = append(res, n.Val)
		ar = ar[1:]
		var t []*Node
		for _, c := range n.Children {
			t = append(t, c)
		}
		ar = append(t, ar...)
	}
	return res
}
