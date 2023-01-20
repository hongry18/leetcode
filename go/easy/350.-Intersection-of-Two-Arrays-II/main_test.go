package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		assert.Equal(t, intersect([]int{1, 1, 2, 2, 2}, []int{2, 2}), []int{2, 2})
	})
}
