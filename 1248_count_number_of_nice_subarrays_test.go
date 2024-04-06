package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberOfSubArrays(t *testing.T) {
	testCases := []struct {
		input  []int
		target int
		want   int
	}{
		{
			[]int{1, 1, 2, 1, 1},
			3,
			2,
		},
		{
			[]int{2, 4, 6},
			1,
			0,
		},
		{
			[]int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2},
			2,
			16,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%v,target:%v", i+1, tc.input, tc.target), func(t *testing.T) {
			output := numberOfSubArrays(tc.input, tc.target)
			assert.Equal(t, tc.want, output)
		})
	}
}
