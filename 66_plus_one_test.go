package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlusOne(t *testing.T) {
	testCases := []struct {
		input []int
		want  []int
	}{
		{
			[]int{1, 2, 3},
			[]int{1, 2, 4},
		},
		{
			[]int{4, 3, 2, 1},
			[]int{4, 3, 2, 2},
		},
		{
			[]int{9},
			[]int{1, 0},
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%v", i+1, tc.input), func(t *testing.T) {
			output := plusOne(tc.input)
			assert.Equal(t, tc.want, output)
		})
	}
}
