package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		assert.Equal(t, fib(2), 1)
		assert.Equal(t, fib(3), 2)
		assert.Equal(t, fib(4), 3)
		assert.Equal(t, fib(20), 6765)
		assert.Equal(t, fib(30), 832040)
	})
}
