package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMedianOfTwoSortedArrays(t *testing.T) {
	testCases := []struct {
		input  []int
		target []int
		want   float64
	}{
		{
			[]int{1, 3},
			[]int{2},
			2,
		},
		{
			[]int{1, 2},
			[]int{3, 4},
			2.5,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%v,input2:%v", i+1, tc.input, tc.target), func(t *testing.T) {
			output := findMedianSortedArrays(tc.input, tc.target)
			assert.Equal(t, tc.want, output)
		})
	}
}
