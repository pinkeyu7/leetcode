package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwapPairs(t *testing.T) {
	testCases := []struct {
		input []int
		want  []int
	}{
		{
			[]int{1, 2, 3, 4},
			[]int{2, 1, 4, 3},
		},
		{
			[]int{},
			[]int{},
		},
		{
			[]int{1},
			[]int{1},
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%v,", i+1, tc.input), func(t *testing.T) {
			input := generateListNode(tc.input)

			output := swapPairs(input)

			list := generateList(output)
			assert.Equal(t, list, tc.want)
		})
	}
}
