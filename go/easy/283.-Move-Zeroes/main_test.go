package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	t.Run("Test1", func(t *testing.T) {
		n := []int{0, 1, 0, 3, 12}
		r := []int{1, 3, 12, 0, 0}
		moveZeroes(n)
		assert.Equal(t, n, r)
	})
}
