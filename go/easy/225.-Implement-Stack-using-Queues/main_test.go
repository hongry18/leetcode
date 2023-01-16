package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	t.Run("push test", func(t *testing.T) {
		x := 5
		s := Constructor()
		s.Push(x)

		assert.Equal(t, s.ar[0], x)
		assert.Equal(t, s.pos, 1)
	})
}
