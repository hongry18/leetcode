package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		assert.Equal(t, intersection([]int{4,9,5}, []int{9,4,9,8,4}), []int{4,9})
	})
}
