package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	assert.Equal(t, romanToInt("III"), 3)
	assert.Equal(t, romanToInt("LVIII"), 58)
	assert.Equal(t, romanToInt("MCMXCIV"), 1994)
}
