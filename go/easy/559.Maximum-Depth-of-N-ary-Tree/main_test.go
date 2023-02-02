package main

import "testing"

func TestXxx(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		r := &Node{Val: 1}
		n1 := &Node{Val: 3}
		r.Children = append(r.Children, n1, &Node{Val: 2}, &Node{Val: 4})
		n1.Children = append(n1.Children, &Node{Val: 5}, &Node{Val: 6})
		maxDepth(r)
	})
}
