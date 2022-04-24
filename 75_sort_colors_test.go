package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortColors(t *testing.T) {
	testCases := []struct {
		input []int
		want  []int
	}{
		{
			[]int{2, 0, 2, 1, 1, 0},
			[]int{0, 0, 1, 1, 2, 2},
		},
		{
			[]int{2, 0, 1},
			[]int{0, 1, 2},
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%v", i+1, tc.input), func(t *testing.T) {
			output := sortColors(tc.input)
			assert.Equal(t, tc.want, output)
		})
	}
}
