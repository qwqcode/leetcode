package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsLetter(t *testing.T) {
	for i := 'a'; i <= 'z'; i++ {
		result := isLetter(i)
		assert.Equal(t, true, result, "Expected true, but got %v", result)
	}

	for i := 'A'; i <= 'Z'; i++ {
		result := isLetter(i)
		assert.Equal(t, true, result, "Expected true, but got %v", result)
	}
}

func TestIsPalindrome(t *testing.T) {
	result := isPalindrome("race a car")
	assert.Equal(t, false, result, "Expected true, but got %v", result)
}
