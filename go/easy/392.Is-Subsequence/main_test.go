package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		assert.Equal(t, isSubsequence("abc", "ahbgdc"), true)
		assert.Equal(t, isSubsequence("axc", "ahbgdc"), false)
		assert.Equal(t, isSubsequence("abcdefgh", "ahbgdc"), false)
		assert.Equal(t, isSubsequence("abc", "ahbgdcd"), true)
	})
}
