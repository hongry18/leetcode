package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		assert.Equal(t, thirdMax([]int{3, 2, 1}), 1)
		assert.Equal(t, thirdMax([]int{1, 2}), 2)
		assert.Equal(t, thirdMax([]int{2, 2, 3, 1}), 1)
	})

	t.Run("test2", func(t *testing.T) {
		t.Log(math.MinInt32)
	})
}
