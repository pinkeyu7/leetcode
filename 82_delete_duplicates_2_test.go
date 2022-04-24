package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteDuplicates2(t *testing.T) {
	testCases := []struct {
		input []int
		want  []int
	}{
		{
			[]int{1, 2, 3, 3, 4, 4, 5},
			[]int{1, 2, 5},
		},
		{
			[]int{1, 1, 1, 2, 3},
			[]int{2, 3},
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%v", i+1, tc.input), func(t *testing.T) {
			list := generateListNode(tc.input)

			output := deleteDuplicates2(list)

			res := generateList(output)
			assert.Equal(t, tc.want, res)
		})
	}
}
