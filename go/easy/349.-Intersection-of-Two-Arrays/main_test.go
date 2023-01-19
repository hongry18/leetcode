package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		// assert.Equal(t, intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}), []int{4, 9})
		// assert.Equal(t, intersection([]int{1, 2, 2, 1}, []int{2, 2}), []int{2})
		assert.Equal(t, intersection([]int{4, 7, 9, 7, 6, 7}, []int{5, 0, 0, 6, 1, 6, 2, 2, 4}), []int{4, 6})
	})
}
