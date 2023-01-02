package mian

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		nums := []int{3, 2, 2, 3}
		expectedNums := []int{2, 2}

		k := removeElement(nums, 3)

		assert.Equal(t, len(expectedNums), k)
		for i := 0; i < k; i++ {
			assert.Equal(t, nums[i], expectedNums[i])
		}
	})
}
