package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		assert.Equal(t, summaryRanges([]int{0, 1, 2, 4, 5, 7}), []string{"0->2", "4->5", "7"})
	})

	t.Run("test2", func(t *testing.T) {
		assert.Equal(t, summaryRanges([]int{0, 2, 4, 5, 6}), []string{"0", "2", "4->6"})
	})

	t.Run("test3", func(t *testing.T) {
		assert.Equal(t, summaryRanges([]int{1}), []string{"1"})
		assert.Equal(t, summaryRanges([]int{1, 2}), []string{"1->2"})
		assert.Equal(t, summaryRanges([]int{1, 3}), []string{"1", "3"})
	})
}
