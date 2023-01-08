package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	assert.Equal(t, singleNumber([]int{2, 2, 1}), 1)
	assert.Equal(t, singleNumber([]int{1}), 1)
	assert.Equal(t, singleNumber([]int{4, 1, 2, 1, 2}), 4)
}
