package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindDuplicate(t *testing.T) {
	testCases := []struct {
		nums []int
		want int
	}{

		{
			[]int{1, 3, 4, 2, 2},
			2,
		},
		{
			[]int{3, 1, 3, 4, 2},
			3,
		},
		{
			[]int{3, 3, 3, 3, 3},
			3,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:nums:%v", i+1, tc.nums), func(t *testing.T) {
			output := findDuplicate(tc.nums)
			assert.Equal(t, output, tc.want)
		})
	}
}
