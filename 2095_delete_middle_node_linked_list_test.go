package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteMiddleNodeLinkedList(t *testing.T) {
	testCases := []struct {
		input []int
		want  []int
	}{
		{
			[]int{1, 3, 4, 7, 1, 2, 6},
			[]int{1, 3, 4, 1, 2, 6},
		},
		{
			[]int{1, 2, 3, 4},
			[]int{1, 2, 4},
		},
		{
			[]int{2, 1},
			[]int{2},
		},
		{
			[]int{1},
			[]int{},
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%v,", i+1, tc.input), func(t *testing.T) {
			input := generateListNode(tc.input)

			output := deleteMiddle(input)

			list := generateList(output)
			assert.Equal(t, tc.want, list)
		})
	}
}
