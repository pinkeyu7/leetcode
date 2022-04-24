package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeKLists(t *testing.T) {
	testCases := []struct {
		input [][]int
		want  []int
	}{
		{
			[][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}},
			[]int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		{
			[][]int{},
			[]int{},
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%v,", i+1, tc.input), func(t *testing.T) {
			lists := make([]*ListNode, 0)
			for _, list := range tc.input {
				lists = append(lists, generateListNode(list))
			}

			output := mergeKLists(lists)
			list := generateList(output)
			assert.Equal(t, list, tc.want)
		})
	}
}
