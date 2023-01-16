package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	t.Run("push", func(t *testing.T) {
		x := 5
		s := Constructor()
		s.Push(x)

		assert.Equal(t, s.ar[0], x)
		assert.Equal(t, s.pos, 0)
	})

	t.Run("push pop", func(t *testing.T) {
		x := 5
		s := Constructor()
		s.Push(x)
		p := s.Pop()

		t.Log(s.ar)
		assert.Equal(t, p, x)
		assert.Equal(t, s.pos, -1)
	})

	t.Run("stack test", func(t *testing.T) {
		s := Constructor()
		s.Push(1)
		s.Push(2)
		t.Log(s.Top())
		t.Log(s.Pop())
		t.Log(s.Empty())
	})
}
