package main

import "testing"

func TestXxx(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		result := sortedArrayToBST([]int{-10, -3, 0, 5, 9})
		q := []*TreeNode{result}

		for len(q) != 0 {
			n := q[0]
			q = q[1:]

			if n == nil {
				continue
			}

			t.Log(n.Val)
			q = append(q, n.Left, n.Right)
		}
	})

	t.Run("test2", func(t *testing.T) {
		result := sortedArrayToBST([]int{1, 3})
		q := []*TreeNode{result}

		for len(q) != 0 {
			n := q[0]
			q = q[1:]

			if n == nil {
				continue
			}

			t.Log(n.Val)
			q = append(q, n.Left, n.Right)
		}
	})

	t.Run("test3", func(t *testing.T) {
		result := sortedArrayToBST([]int{1, 3, 5, 7})
		q := []*TreeNode{result}

		for len(q) != 0 {
			n := q[0]
			q = q[1:]

			if n == nil {
				continue
			}

			t.Log(n.Val)
			q = append(q, n.Left, n.Right)
		}
	})

	t.Run("test4", func(t *testing.T) {
		result := sortedArrayToBST([]int{-1, 0, 1, 2})
		q := []*TreeNode{result}

		for len(q) != 0 {
			n := q[0]
			q = q[1:]

			if n == nil {
				continue
			}

			t.Log(n.Val)
			q = append(q, n.Left, n.Right)
		}
	})

	t.Run("test5", func(t *testing.T) {
		result := sortedArrayToBST([]int{0, 1, 2, 3, 4, 5})
		q := []*TreeNode{result}

		for len(q) != 0 {
			n := q[0]
			q = q[1:]

			if n == nil {
				continue
			}

			t.Log(n.Val)
			q = append(q, n.Left, n.Right)
		}
	})
}
