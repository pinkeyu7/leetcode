package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	testCases := []struct {
		nums1 []int
		nums2 []int
		want  []int
	}{
		{
			[]int{2, 4, 3},
			[]int{5, 6, 4},
			[]int{7, 0, 8},
		},
		{
			[]int{0},
			[]int{0},
			[]int{0},
		},
		{
			[]int{9, 9, 9, 9, 9, 9, 9},
			[]int{9, 9, 9, 9},
			[]int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:nums1:%v,nums2:%v,", i+1, tc.nums1, tc.nums2), func(t *testing.T) {
			l1 := generateListNode(tc.nums1)
			l2 := generateListNode(tc.nums2)

			output := addTwoNumbers(l1, l2)

			list := generateList(output)
			assert.Equal(t, list, tc.want)
		})
	}
}
