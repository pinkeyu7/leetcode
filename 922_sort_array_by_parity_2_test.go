package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortArrayByParityII(t *testing.T) {
	testCases := []struct {
		input []int
		want  []int
	}{
		{
			[]int{4, 2, 5, 7},
			[]int{4, 5, 2, 7},
		},
		{
			[]int{2, 3},
			[]int{2, 3},
		},
		{
			[]int{2, 3, 1, 1, 4, 0, 0, 4, 3, 3},
			[]int{2, 3, 0, 1, 4, 1, 0, 3, 4, 3},
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%v", i+1, tc.input), func(t *testing.T) {
			output := sortArrayByParityII(tc.input)
			assert.Equal(t, tc.want, output)
		})
	}
}
