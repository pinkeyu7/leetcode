package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncreasingTripletSubsequence(t *testing.T) {
	testCases := []struct {
		input []int
		want  bool
	}{
		{
			[]int{1, 2, 3, 4, 5},
			true,
		},
		{
			[]int{5, 4, 3, 2, 1},
			false,
		},
		{
			[]int{2, 1, 5, 0, 4, 6},
			true,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%v", i+1, tc.input), func(t *testing.T) {
			output := increasingTriplet(tc.input)
			assert.Equal(t, tc.want, output)
		})
	}
}
