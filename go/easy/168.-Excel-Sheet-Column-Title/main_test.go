package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		assert.Equal(t, convertToTitle(1), "A")
		assert.Equal(t, convertToTitle(26), "Z")
	})

	t.Run("test2", func(t *testing.T) {
		assert.Equal(t, convertToTitle(28), "AB")
		assert.Equal(t, convertToTitle(701), "ZY")
	})

	t.Run("test3", func(t *testing.T) {
		assert.Equal(t, convertToTitle(703), "AAA")
	})
}
