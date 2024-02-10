package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestDivisibleSubset(t *testing.T) {
	testCases := []struct {
		input []int
		want  []int
	}{
		{
			[]int{1},
			[]int{1},
		},
		{
			[]int{1, 2, 3},
			[]int{1, 2},
		},
		{
			[]int{1, 2, 4, 8},
			[]int{1, 2, 4, 8},
		},
		{
			[]int{4, 8, 10, 240},
			[]int{4, 8, 240},
		},
		{
			[]int{5, 9, 18, 54, 108, 540, 90, 180, 360, 720},
			[]int{9, 18, 90, 180, 360, 720},
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%v", i+1, tc.input), func(t *testing.T) {
			output := largestDivisibleSubset(tc.input)
			assert.Equal(t, tc.want, output)
		})
	}
}
