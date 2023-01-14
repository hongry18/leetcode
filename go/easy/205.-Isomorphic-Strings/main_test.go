package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		assert.Equal(t, isIsomorphic("egg", "add"), true)
		assert.Equal(t, isIsomorphic("foo", "bar"), false)
		assert.Equal(t, isIsomorphic("paper", "title"), true)
	})
}
