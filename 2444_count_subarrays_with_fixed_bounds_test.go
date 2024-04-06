package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountSubArrays(t *testing.T) {
	testCases := []struct {
		input []int
		minK  int
		maxK  int
		want  int64
	}{
		{
			[]int{1, 3, 5, 2, 7, 5},
			1,
			5,
			2,
		},
		{
			[]int{1, 1, 1, 1},
			1,
			1,
			10,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%v,minK:%v,maxK:%v", i+1, tc.input, tc.minK, tc.maxK), func(t *testing.T) {
			output := countSubArrays(tc.input, tc.minK, tc.maxK)
			assert.Equal(t, tc.want, output)
		})
	}
}
