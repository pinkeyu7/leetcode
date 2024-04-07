package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestConsecutive(t *testing.T) {
	testCases := []struct {
		nums []int
		want int
	}{
		{
			[]int{100, 4, 200, 1, 3, 2},
			4,
		},
		{
			[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			9,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:nums:%v", i+1, tc.nums), func(t *testing.T) {
			output := longestConsecutive(tc.nums)
			assert.Equal(t, output, tc.want)
		})
	}
}
