package main

type Node struct {
	Val      int
	Children []*Node
}

func travel(n *Node, res *[]int) {
	if n == nil {
		return
	}
	for _, c := range n.Children {
		travel(c, res)
	}
	*res = append(*res, n.Val)
}

func reverse(root *Node) []int {
	if root == nil {
		return []int{}
	}
	var res []int
	var ar = []*Node{root}
	for len(ar) > 0 {
		l := ar[len(ar)-1]
		ar = ar[:len(ar)-1]
		res = append(res, l.Val)
		ar = append(ar, l.Children...)
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	return res
}

func postorder(root *Node) []int {
	var res []int
	if root == nil {
		return res
	}

	var ar = []*Node{root}
	for len(ar) > 0 {
		cur := ar[len(ar)-1]
		if len(cur.Children) == 0 {
			ar = ar[:len(ar)-1]
			res = append(res, cur.Val)
		} else {
			for i := len(cur.Children) - 1; i >= 0; i-- {
				ar = append(ar, cur.Children[i])
			}
			cur.Children = []*Node{}
		}
	}

	return res
}
