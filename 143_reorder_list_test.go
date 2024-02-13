package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReorderList(t *testing.T) {
	testCases := []struct {
		input []int
		want  []int
	}{
		{
			[]int{1, 2, 3, 4},
			[]int{1, 4, 2, 3},
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{1, 5, 2, 4, 3},
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%v,", i+1, tc.input), func(t *testing.T) {
			input := generateListNode(tc.input)

			output := reorderList(input)

			list := generateList(output)
			assert.Equal(t, tc.want, list)
		})
	}
}
