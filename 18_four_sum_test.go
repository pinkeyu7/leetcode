package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFourSum(t *testing.T) {
	testCases := []struct {
		nums   []int
		target int
		want   [][]int
	}{

		{
			[]int{1, 0, -1, 0, -2, 2},
			0,
			[][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}},
		},
		{
			[]int{2, 2, 2, 2, 2},
			8,
			[][]int{{2, 2, 2, 2}},
		},
		{
			[]int{-3, -1, 0, 2, 4, 5},
			0,
			[][]int{{-3, -1, 0, 4}},
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:nums:%v,target:%d,", i+1, tc.nums, tc.target), func(t *testing.T) {
			output := fourSum(tc.nums, tc.target)
			assert.Equal(t, output, tc.want)
		})
	}
}
