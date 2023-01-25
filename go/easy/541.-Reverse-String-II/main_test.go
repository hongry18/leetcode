package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		assert.Equal(t, reverseStr("abcdefghijklm", 5), "edcbafghijmlk")
		assert.Equal(t, reverseStr("abcdefg", 2), "bacdfeg")
		assert.Equal(t, reverseStr("abcd", 2), "bacd")
	})
}
