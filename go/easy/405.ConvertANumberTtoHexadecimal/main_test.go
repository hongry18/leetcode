package main

import "testing"

func TestXxx(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		toHex(26)
		toHex(-1)
	})
}
