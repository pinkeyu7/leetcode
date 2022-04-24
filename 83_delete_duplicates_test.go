package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteDuplicates(t *testing.T) {
	testCases := []struct {
		input []int
		want  []int
	}{
		{
			[]int{1, 1, 2},
			[]int{1, 2},
		},
		{
			[]int{1, 1, 2, 3, 3},
			[]int{1, 2, 3},
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%v", i+1, tc.input), func(t *testing.T) {
			list := generateListNode(tc.input)

			output := deleteDuplicates(list)

			res := generateList(output)
			assert.Equal(t, tc.want, res)
		})
	}
}
