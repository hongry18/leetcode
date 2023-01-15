package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		assert.Equal(t, containsDuplicate([]int{1, 2, 3, 1}), true)
		assert.Equal(t, containsDuplicate([]int{1, 2, 3, 4}), false)
		assert.Equal(t, containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}), true)
	})
}
