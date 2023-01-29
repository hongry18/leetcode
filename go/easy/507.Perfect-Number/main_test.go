package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		assert.Equal(t, checkPerfectNumber(28), true)
		assert.Equal(t, checkPerfectNumber(7), false)
	})
}
