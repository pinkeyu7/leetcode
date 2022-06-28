package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseBetween(t *testing.T) {
	testCases := []struct {
		input []int
		left  int
		right int
		want  []int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			2,
			4,

			[]int{1, 4, 3, 2, 5},
		},
		{
			[]int{5},
			1,
			1,
			[]int{5},
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%v,left:%d,right:%d", i+1, tc.input, tc.left, tc.right), func(t *testing.T) {
			input := generateListNode(tc.input)

			output := reverseBetween(input, tc.left, tc.right)

			list := generateList(output)
			assert.Equal(t, list, tc.want)
		})
	}
}
