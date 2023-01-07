package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	assert.Equal(t, isPalindrome("A man, a plan, a canal: Panama"), true)
	assert.Equal(t, isPalindrome("race a car"), false)
	assert.Equal(t, isPalindrome(" "), true)
}
