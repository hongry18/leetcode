package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	assert.Equal(t, maxProfit([]int{7, 1, 5, 3, 6, 4}), 5)
	assert.Equal(t, maxProfit([]int{7, 6, 4, 3, 1}), 0)
	assert.Equal(t, maxProfit([]int{2, 4, 1}), 2)
}
