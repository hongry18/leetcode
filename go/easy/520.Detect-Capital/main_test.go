package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		assert.Equal(t, detectCapitalUse("leetcode"), true)
		assert.Equal(t, detectCapitalUse("USA"), true)
		assert.Equal(t, detectCapitalUse("FlaG"), false)
		assert.Equal(t, detectCapitalUse("Leetcode"), true)
		assert.Equal(t, detectCapitalUse("LEetcode"), false)
	})
}
