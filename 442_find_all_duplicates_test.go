package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindDuplicates(t *testing.T) {
	testCases := []struct {
		nums []int
		want []int
	}{

		{
			[]int{4, 3, 2, 7, 8, 2, 3, 1},
			[]int{2, 3},
		},
		{
			[]int{1, 1, 2},
			[]int{1},
		},
		{
			[]int{1},
			[]int{},
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:nums:%v", i+1, tc.nums), func(t *testing.T) {
			output := findDuplicates(tc.nums)
			assert.Equal(t, output, tc.want)
		})
	}
}
