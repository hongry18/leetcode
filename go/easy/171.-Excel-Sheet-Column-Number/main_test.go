package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		assert.Equal(t, titleToNumber("A"), 1)
		assert.Equal(t, titleToNumber("Z"), 26)
		assert.Equal(t, titleToNumber("AB"), 28)
		assert.Equal(t, titleToNumber("ZY"), 701)
	})
}
