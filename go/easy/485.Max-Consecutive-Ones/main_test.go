package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		assert.Equal(t, findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}), 3)
		assert.Equal(t, findMaxConsecutiveOnes([]int{1, 0, 1, 1, 0, 1}), 2)
	})

	t.Run("test2", func(t *testing.T) {
		assert.Equal(t, findMaxConsecutiveOnes([]int{1, 1, 0, 1}), 2)
	})
}
