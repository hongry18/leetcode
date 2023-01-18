package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		nums := []int{-2, 0, 3, -5, 2, -1}

		obj := Constructor(nums)
		assert.Equal(t, obj.SumRange(0, 2), 1)
		assert.Equal(t, obj.SumRange(2, 5), -1)
		assert.Equal(t, obj.SumRange(0, 5), -3)
	})
}
