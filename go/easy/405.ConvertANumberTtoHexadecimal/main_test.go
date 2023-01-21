package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		assert.Equal(t, toHex(26), "1a")
		assert.Equal(t, toHex(-1), "ffffffff")
	})
}
