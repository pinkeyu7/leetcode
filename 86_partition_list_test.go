package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartitionList(t *testing.T) {
	testCases := []struct {
		input  []int
		target int
		want   []int
	}{
		{
			[]int{1, 4, 3, 2, 5, 2},
			3,
			[]int{1, 2, 2, 4, 3, 5},
		},
		{
			[]int{2, 1},
			2,
			[]int{1, 2},
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%v,target:%v,", i+1, tc.input, tc.target), func(t *testing.T) {
			input := generateListNode(tc.input)

			output := partition(input, tc.target)

			list := generateList(output)
			assert.Equal(t, tc.want, list)
		})
	}
}
