package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		assert.Equal(t, canConstruct("aab", "baa"), true)
		assert.Equal(t, canConstruct("aa", "aab"), true)
		assert.Equal(t, canConstruct("a", "b"), false)
		assert.Equal(t, canConstruct("aa", "ab"), false)
	})

	t.Run("test2", func(t *testing.T) {
		assert.Equal(t, canConstruct("fihjjjjei", "hjibagacbhadfaefdjaeaebgi"), false)
	})

	t.Run("test3", func(t *testing.T) {
		assert.Equal(t, canConstruct("bg", "gfbbgbja"), true)
	})
}
