package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		nums   []int
		target int
		want   []int
	}{
		{
			[]int{2, 7, 11, 15},
			9,
			[]int{0, 1},
		},
		{
			[]int{3, 2, 4},
			6,
			[]int{1, 2},
		},
		{
			[]int{3, 3},
			6,
			[]int{0, 1},
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:nums:%v,target:%d,", i+1, tc.nums, tc.target), func(t *testing.T) {
			output := twoSum(tc.nums, tc.target)
			assert.Equal(t, output, tc.want)
		})
	}
}
