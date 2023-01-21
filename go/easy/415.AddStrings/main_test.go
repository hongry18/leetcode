package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	t.Run("test1", func(t *testing.T) { assert.Equal(t, addStrings("11", "123"), "134") })

	t.Run("test2", func(t *testing.T) { assert.Equal(t, addStrings("456", "77"), "533") })
}
