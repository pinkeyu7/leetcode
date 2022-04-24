package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniquePathsWithObstacles(t *testing.T) {
	testCases := []struct {
		input [][]int
		want  int
	}{
		{
			[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
			2,
		},
		{
			[][]int{{0, 1}, {0, 0}},
			1,
		},
		{
			[][]int{{0, 0}, {0, 1}},
			0,
		},
		{
			[][]int{{1, 0}},
			0,
		},
		{
			[][]int{{0, 1}, {1, 0}},
			0,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%v", i+1, tc.input), func(t *testing.T) {
			output := uniquePathsWithObstacles(tc.input)
			assert.Equal(t, tc.want, output)
		})
	}
}
