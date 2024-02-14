package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumOfAbsoluteDifferencesInSortedArray(t *testing.T) {
	testCases := []struct {
		input []int
		want  []int
	}{
		{
			[]int{2, 3, 5},
			[]int{4, 3, 5},
		},
		{
			[]int{1, 4, 6, 8, 10},
			[]int{24, 15, 13, 15, 21},
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%v", i+1, tc.input), func(t *testing.T) {
			output := getSumAbsoluteDifferences(tc.input)
			assert.Equal(t, tc.want, output)
		})
	}
}
