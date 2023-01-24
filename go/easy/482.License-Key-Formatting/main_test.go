package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		assert.Equal(t, licenseKeyFormatting("5f3z-2e-9-w", 4), "5F3Z-2E9W")
		assert.Equal(t, licenseKeyFormatting("2-5g-3-J", 2), "2-5G-3J")
	})
}
