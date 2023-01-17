package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	t.Run("tes1", func(t *testing.T) {
		assert.Equal(t, addDigits(38), 2)
	})
}
