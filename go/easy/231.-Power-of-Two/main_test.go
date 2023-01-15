package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		assert.Equal(t, isPowerOfTwo(1), true)
		assert.Equal(t, isPowerOfTwo(2), true)
		assert.Equal(t, isPowerOfTwo(4), true)
	})
}
