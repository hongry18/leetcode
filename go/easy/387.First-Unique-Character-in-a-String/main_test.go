package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		assert.Equal(t, firstUniqChar("leetcode"), 0)
		assert.Equal(t, firstUniqChar("loveleetcode"), 2)
		assert.Equal(t, firstUniqChar("aabb"), -1)
	})
}
