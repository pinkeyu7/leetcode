package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMostFrequentEvenElement(t *testing.T) {
	testCases := []struct {
		input []int
		want  int
	}{
		{
			[]int{0, 1, 2, 2, 4, 4, 1},
			2,
		},
		{
			[]int{4, 4, 4, 9, 2, 4},
			4,
		},
		{
			[]int{29, 47, 21, 41, 13, 37, 25, 7},
			-1,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%v", i+1, tc.input), func(t *testing.T) {
			output := mostFrequentEven(tc.input)
			assert.Equal(t, tc.want, output)
		})
	}
}
