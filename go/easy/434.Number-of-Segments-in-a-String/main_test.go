package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		assert.Equal(t, countSegments("Hello, my name is John"), 5)
		assert.Equal(t, countSegments("  Hello,-my name is John   "), 4)
		// assert.Equal(t, countSegments(""), 5)
		// assert.Equal(t, countSegments(""), 5)
	})
}
