package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		assert.Equal(t, wordPattern("abba", "dog cat cat dog"), true)
		assert.Equal(t, wordPattern("aaaa", "dog cat cat dog"), false)
		assert.Equal(t, wordPattern("b", "dog"), true)
	})

	t.Run("test2", func(t *testing.T) {
		assert.Equal(t, wordPattern("abba", "dog cat cat fish"), false)
	})
}
