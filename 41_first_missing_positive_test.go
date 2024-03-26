package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstMissingPositive(t *testing.T) {
	testCases := []struct {
		nums []int
		want int
	}{

		{
			[]int{1, 2, 0},
			3,
		},
		{
			[]int{3, 4, -1, 1},
			2,
		},
		{
			[]int{7, 8, 9, 11, 12},
			1,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:nums:%v", i+1, tc.nums), func(t *testing.T) {
			output := firstMissingPositive(tc.nums)
			assert.Equal(t, output, tc.want)
		})
	}
}
