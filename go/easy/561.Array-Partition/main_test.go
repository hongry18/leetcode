package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		assert.Equal(t, arrayPairSum([]int{6, 2, 6, 5, 1, 2}), 9)
	})
}
