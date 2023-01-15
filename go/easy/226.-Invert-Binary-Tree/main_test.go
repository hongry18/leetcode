package main

import "testing"

func TestXxx(t *testing.T) {
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 7}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 9}

	r := invertTree(root)

	ar := []*TreeNode{r}
	for len(ar) != 0 {
		n := ar[0]
		ar = ar[1:]
		if n == nil {
			continue
		}
		t.Log(n.Val)
		ar = append(ar, n.Left, n.Right)
	}
}
