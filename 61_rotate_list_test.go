package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotateRight(t *testing.T) {
	testCases := []struct {
		input  []int
		target int
		want   []int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			2,
			[]int{4, 5, 1, 2, 3},
		},
		{
			[]int{0, 1, 2},
			4,
			[]int{2, 0, 1},
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%v,target:%d,", i+1, tc.input, tc.target), func(t *testing.T) {
			input := generateListNode(tc.input)

			output := rotateRight(input, tc.target)

			list := generateList(output)
			assert.Equal(t, tc.want, list)
		})
	}
}
