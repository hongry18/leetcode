package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	assert.Equal(t, addBinary("11", "1"), "100")
	assert.Equal(t, addBinary("1010", "1011"), "10101")
}
