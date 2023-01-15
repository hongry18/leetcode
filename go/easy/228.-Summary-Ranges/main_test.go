package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		assert.Equal(t, summaryRanges([]int{0, 1, 2, 4, 5, 7}), []string{"0->2", "4->5", "7"})
	})
}
