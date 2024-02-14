package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartitionArrayAccordingToGivenPivot(t *testing.T) {
	testCases := []struct {
		input  []int
		target int
		want   []int
	}{
		{
			[]int{9, 12, 5, 10, 14, 3, 10},
			10,
			[]int{9, 5, 3, 10, 10, 12, 14},
		},
		{
			[]int{-3, 4, 3, 2},
			2,
			[]int{-3, 2, 4, 3},
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%v,target:%v", i+1, tc.input, tc.target), func(t *testing.T) {
			output := pivotArray(tc.input, tc.target)
			assert.Equal(t, tc.want, output)
		})
	}
}
