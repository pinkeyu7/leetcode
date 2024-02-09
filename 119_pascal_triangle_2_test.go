package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPascalTriangle2(t *testing.T) {
	testCases := []struct {
		target int
		want   []int
	}{
		{
			1,
			[]int{1, 1},
		},
		{
			2,
			[]int{1, 2, 1},
		},
		{
			4,
			[]int{1, 4, 6, 4, 1},
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:target:%d,", i+1, tc.target), func(t *testing.T) {
			output := getRow(tc.target)
			assert.Equal(t, tc.want, output)
		})
	}
}
