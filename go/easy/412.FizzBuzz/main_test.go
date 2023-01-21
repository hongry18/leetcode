package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		assert.Equal(t, fizzBuzz(3), []string{"1", "2", "Fizz"})
		assert.Equal(t, fizzBuzz(5), []string{"1", "2", "Fizz", "4", "Buzz"})
		assert.Equal(t, fizzBuzz(15), []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"})
	})
}
