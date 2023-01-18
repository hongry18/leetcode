package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		assert.Equal(t, isPowerOfFour(16), true)
	})

	t.Run("test2", func(t *testing.T) {
		assert.Equal(t, isPowerOfFour(8), false)
	})
}
