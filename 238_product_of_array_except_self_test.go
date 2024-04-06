package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductExceptSelf(t *testing.T) {
	testCases := []struct {
		nums []int
		want []int
	}{
		{
			[]int{1, 2, 3, 4},
			[]int{24, 12, 8, 6},
		},
		{
			[]int{-1, 1, 0, -3, 3},
			[]int{0, 0, 9, 0, 0},
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:nums:%v,", i+1, tc.nums), func(t *testing.T) {
			output := productExceptSelf(tc.nums)
			assert.Equal(t, output, tc.want)
		})
	}
}
