package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		assert.Equal(t, containsNearbyDuplicate([]int{1, 2, 3, 1}, 3), true)
	})

	t.Run("test2", func(t *testing.T) {
		assert.Equal(t, containsNearbyDuplicate([]int{1, 0, 1, 1}, 1), true)
	})

	t.Run("test3", func(t *testing.T) {
		assert.Equal(t, containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2), false)
	})

	t.Run("test4", func(t *testing.T) {
		assert.Equal(t, containsNearbyDuplicate([]int{1, 1}, 2), true)
	})
}
