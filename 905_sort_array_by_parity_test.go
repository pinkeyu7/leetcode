package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortArrayByParity(t *testing.T) {
	testCases := []struct {
		input []int
		want  []int
	}{
		{
			[]int{3, 1, 2, 4},
			[]int{4, 2, 1, 3},
		},
		{
			[]int{0},
			[]int{0},
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%v", i+1, tc.input), func(t *testing.T) {
			output := sortArrayByParity(tc.input)
			assert.Equal(t, tc.want, output)
		})
	}
}
