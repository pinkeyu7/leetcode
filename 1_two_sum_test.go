package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9

	output := twoSum(nums, target)
	assert.Equal(t, output, []int{0, 1})
}
