package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		assert.Equal(t, findRelativeRanks([]int{10, 3, 8, 9, 4}), []string{"Gold Medal", "5", "Bronze Medal", "Silver Medal", "4"})
	})
}
