package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		assert.Equal(t, checkRecord("PPALLP"), true)
		assert.Equal(t, checkRecord("PPALLL"), false)
	})
}
