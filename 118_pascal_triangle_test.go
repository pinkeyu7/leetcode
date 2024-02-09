package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPascalTriangle(t *testing.T) {
	testCases := []struct {
		target int
		want   [][]int
	}{
		{
			1,
			[][]int{{1}},
		},
		{
			2,
			[][]int{{1}, {1, 1}},
		},
		{
			5,
			[][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}},
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:target:%d,", i+1, tc.target), func(t *testing.T) {
			output := generate(tc.target)
			assert.Equal(t, tc.want, output)
		})
	}
}
