package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinPathSum(t *testing.T) {
	testCases := []struct {
		input [][]int
		want  int
	}{
		{
			[][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}},
			7,
		},
		{
			[][]int{{1, 2, 3}, {4, 5, 6}},
			12,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%v", i+1, tc.input), func(t *testing.T) {
			output := minPathSum(tc.input)
			assert.Equal(t, tc.want, output)
		})
	}
}
