package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	assert.Equal(t, plusOne([]int{1, 2, 3}), []int{1, 2, 4})
	assert.Equal(t, plusOne([]int{4, 3, 2, 1}), []int{4, 3, 2, 2})
	assert.Equal(t, plusOne([]int{9}), []int{1, 0})
}
