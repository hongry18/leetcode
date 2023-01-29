package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		assert.Equal(t, nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}), []int{-1, 3, -1})
	})
}
