package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		assert.Equal(t, isPowerOfThree(27), true)
		assert.Equal(t, isPowerOfThree(1), true)
		assert.Equal(t, isPowerOfThree(8), false)
		assert.Equal(t, isPowerOfThree(81), true)
		assert.Equal(t, isPowerOfThree(0), false)
		assert.Equal(t, isPowerOfThree(-1), false)
	})

	t.Run("test2", func(t *testing.T) {
		assert.Equal(t, isPowerOfThree(21), false)
	})
}
