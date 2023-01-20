package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		assert.Equal(t, findTheDifference("abcd", "abcde"), byte('e'))
		assert.Equal(t, findTheDifference("", "y"), byte('y'))
	})
}
