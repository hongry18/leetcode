package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		r := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  2,
					Next: &ListNode{Val: 1},
				},
			},
		}

		assert.Equal(t, isPalindrome(r), true)
	})
}
