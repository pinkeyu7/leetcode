package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopKFrequentElements(t *testing.T) {
	testCases := []struct {
		input  []int
		target int
		want   []int
	}{
		{
			[]int{1, 1, 1, 2, 2, 3},
			2,
			[]int{1, 2},
		},
		{
			[]int{1},
			1,
			[]int{1},
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%v,target:%v", i+1, tc.input, tc.target), func(t *testing.T) {
			output := topKFrequent(tc.input, tc.target)
			assert.Equal(t, tc.want, output)
		})
	}
}
