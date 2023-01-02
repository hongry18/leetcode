package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClimbStairs(t *testing.T) {
	t.Log(assert.Equal(t, climbStairs(2), 2))
	t.Log(assert.Equal(t, climbStairs(3), 3))
	t.Log(assert.Equal(t, climbStairs(4), 5))
	t.Log(assert.Equal(t, climbStairs(5), 8))
}
