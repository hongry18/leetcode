package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		nums := []int{1, 1, 2}
		expectedNums := []int{1, 2}

		k := removeDuplicates(nums)

		assert.Equal(t, len(expectedNums), k)
		for i := 0; i < k; i++ {
			assert.Equal(t, nums[i], expectedNums[i])
		}
	})
}
