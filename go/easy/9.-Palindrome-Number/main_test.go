package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	assert.Equal(t, isPalindrome(121), true)
	assert.Equal(t, isPalindrome(-121), false)
	assert.Equal(t, isPalindrome(10), false)
}
