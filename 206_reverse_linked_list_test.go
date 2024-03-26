package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseList(t *testing.T) {
	testCases := []struct {
		input []int
		want  []int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			[]int{5, 4, 3, 2, 1},
		},
		{
			[]int{1, 2},
			[]int{2, 1},
		},
		{
			[]int{},
			[]int{},
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%v,", i+1, tc.input), func(t *testing.T) {
			input := generateListNode(tc.input)

			output := reverseList(input)

			list := generateList(output)
			assert.Equal(t, tc.want, list)
		})
	}
}
