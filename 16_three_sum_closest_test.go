package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeSumClosest(t *testing.T) {
	testCases := []struct {
		nums   []int
		target int
		want   int
	}{

		{
			[]int{-1, 2, 1, -4},
			1,
			2,
		},
		{
			[]int{0, 0, 0},
			1,
			0,
		},
		{
			[]int{0, 1, 2},
			0,
			3,
		},
		{
			[]int{-2, -1, 1, 4},
			0,
			1,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:nums:%v,target:%d,", i+1, tc.nums, tc.target), func(t *testing.T) {
			output := threeSumClosest(tc.nums, tc.target)
			assert.Equal(t, tc.want, output)
		})
	}
}
