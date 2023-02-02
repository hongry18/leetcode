package main

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	var cnt int

	ar := []*Node{root}
	for len(ar) != 0 {
		cnt++
		q2 := []*Node{}
		for _, n := range ar {
			q2 = append(q2, n.Children...)
		}
		ar = q2
	}
	return cnt
}
