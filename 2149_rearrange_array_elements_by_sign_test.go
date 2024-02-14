package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRearrangeArrayElementsBySign(t *testing.T) {
	testCases := []struct {
		input []int
		want  []int
	}{
		{
			[]int{3, 1, -2, -5, 2, -4},
			[]int{3, -2, 1, -5, 2, -4},
		},
		{
			[]int{-1, 1},
			[]int{1, -1},
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%v", i+1, tc.input), func(t *testing.T) {
			output := rearrangeArray(tc.input)
			assert.Equal(t, tc.want, output)
		})
	}
}
