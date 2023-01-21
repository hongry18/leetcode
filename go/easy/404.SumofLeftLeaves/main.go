package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	Node *TreeNode
	Pos  bool
}

func sumOfLeftLeaves(root *TreeNode) int {
	var ar = []Node{{Node: root}}

	var sum int
	for len(ar) != 0 {
		n := ar[0]
		ar = ar[1:]
		if n.Node == nil {
			continue
		}

		if n.Node.Left == nil && n.Node.Right == nil && n.Pos {
			sum += n.Node.Val
		}

		ar = append(ar, Node{Node: n.Node.Left, Pos: true}, Node{Node: n.Node.Right, Pos: false})
	}

	return sum
}
