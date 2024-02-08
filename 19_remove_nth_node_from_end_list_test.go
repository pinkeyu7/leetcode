package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveNthFromEnd(t *testing.T) {
	testCases := []struct {
		input  []int
		target int
		want   []int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			2,
			[]int{1, 2, 3, 5},
		},
		{
			[]int{1},
			1,
			[]int{},
		},
		{
			[]int{1, 2},
			1,
			[]int{1},
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%v,", i+1, tc.input), func(t *testing.T) {
			inputList := generateListNode(tc.input)
			output := removeNthFromEnd(inputList, tc.target)
			outputList := generateList(output)
			assert.Equal(t, tc.want, outputList)
		})
	}
}
