package main

import "testing"

func TestXxx(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		q := Constructor()
		q.Push(1)
		q.Push(2)
		t.Log(q.Peek())
		t.Log(q.Pop())
		q.Push(3)
		t.Log(q.Pop())
		t.Log(q.Pop())
		t.Log(q.Empty())
	})
}
