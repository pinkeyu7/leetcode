package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMissingNumber(t *testing.T) {
	testCases := []struct {
		nums []int
		want int
	}{

		{
			[]int{3, 0, 1},
			2,
		},
		{
			[]int{0, 1},
			2,
		},
		{
			[]int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			8,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:nums:%v", i+1, tc.nums), func(t *testing.T) {
			output := missingNumber(tc.nums)
			assert.Equal(t, output, tc.want)
		})
	}
}
