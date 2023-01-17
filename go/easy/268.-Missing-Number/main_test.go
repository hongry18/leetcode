package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		assert.Equal(t, missingNumber([]int{3, 0, 1}), 2)
		assert.Equal(t, missingNumber([]int{0, 1}), 2)
		assert.Equal(t, missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}), 8)
	})
}
