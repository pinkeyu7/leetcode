package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsDuplicate(t *testing.T) {
	testCases := []struct {
		nums []int
		want bool
	}{
		{
			[]int{1, 2, 3, 1},
			true,
		},
		{
			[]int{1, 2, 3, 4},
			false,
		},
		{
			[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			true,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:nums:%v,", i+1, tc.nums), func(t *testing.T) {
			output := containsDuplicate(tc.nums)
			assert.Equal(t, output, tc.want)
		})
	}
}
