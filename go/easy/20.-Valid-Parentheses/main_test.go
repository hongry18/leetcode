package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	assert.Equal(t, isValid("()"), true)
	assert.Equal(t, isValid("()[]{}"), true)
	assert.Equal(t, isValid("(]"), false)
}
