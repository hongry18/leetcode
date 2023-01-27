package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		assert.Equal(t, convertToBase7(100), "202")
		assert.Equal(t, convertToBase7(-7), "-10")
	})
}
