package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestIncreasingSubsequence(t *testing.T) {
	testCases := []struct {
		input []int
		want  int
	}{
		{
			[]int{10, 9, 2, 5, 3, 7, 101, 18},
			4,
		},
		{
			[]int{0, 1, 0, 3, 2, 3},
			4,
		},
		{
			[]int{7, 7, 7, 7, 7, 7, 7},
			1,
		},
		{
			[]int{1, 3, 6, 7, 9, 4, 10, 5, 6},
			6,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%v", i+1, tc.input), func(t *testing.T) {
			output := lengthOfLIS(tc.input)
			assert.Equal(t, tc.want, output)
		})
	}
}
