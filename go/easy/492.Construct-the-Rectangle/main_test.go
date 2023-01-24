package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		assert.Equal(t, constructRectangle(4), []int{2, 2})
		assert.Equal(t, constructRectangle(37), []int{37, 1})
		assert.Equal(t, constructRectangle(122122), []int{427, 286})
	})
}
