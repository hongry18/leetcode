package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		assert.Equal(t, longestPalindrome("abccccdd"), 7)
		assert.Equal(t, longestPalindrome("aaaAaaaa"), 7)
		// t.Log('a', 'A')
	})
}
