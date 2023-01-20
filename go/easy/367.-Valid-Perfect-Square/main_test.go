package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		assert.Equal(t, isPerfectSquare(16), true)
		assert.Equal(t, isPerfectSquare(14), false)
		assert.Equal(t, isPerfectSquare(5), false)
	})
}
