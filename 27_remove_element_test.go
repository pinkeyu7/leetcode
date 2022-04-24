package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveElement(t *testing.T) {
	testCases := []struct {
		input  []int
		target int
		want   int
	}{
		{
			[]int{3, 2, 2, 3},
			3,
			2,
		},
		{
			[]int{0, 1, 2, 2, 3, 0, 4, 2},
			2,
			5,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%v,target:%d,", i+1, tc.input, tc.target), func(t *testing.T) {
			output := removeElement(tc.input, tc.target)
			assert.Equal(t, tc.want, output)
		})
	}
}
