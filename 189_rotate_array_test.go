package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	testCases := []struct {
		input  []int
		target int
		want   []int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6, 7},
			3,
			[]int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			[]int{-1, -100, 3, 99},
			2,
			[]int{3, 99, -1, -100},
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%v,target:%d", i+1, tc.input, tc.target), func(t *testing.T) {
			output := rotate(tc.input, tc.target)
			assert.Equal(t, tc.want, output)
		})
	}
}
