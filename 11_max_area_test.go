package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxArea(t *testing.T) {
	testCases := []struct {
		input []int
		want  int
	}{
		{
			[]int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			49,
		},
		{
			[]int{1, 1},
			1,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:%v", i+1, tc.input), func(t *testing.T) {
			output := maxArea(tc.input)
			assert.Equal(t, output, tc.want)
		})
	}
}
