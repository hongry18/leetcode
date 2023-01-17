package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	t.Run("test2", func(t *testing.T) {
		assert.Equal(t, isUgly(6), true)
		assert.Equal(t, isUgly(1), true)
	})

	t.Run("test1", func(t *testing.T) {
		assert.Equal(t, isUgly(14), false)
	})
}
