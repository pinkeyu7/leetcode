package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermuteUnique(t *testing.T) {
	testCases := []struct {
		input []int
		want  [][]int
	}{
		{
			[]int{1, 2, 3},
			[][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
		{
			[]int{1, 1, 2},
			[][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}},
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%v", i+1, tc.input), func(t *testing.T) {
			output := permuteUnique(tc.input)
			assert.Equal(t, sortInt2D(tc.want), sortInt2D(output))
		})
	}
}
