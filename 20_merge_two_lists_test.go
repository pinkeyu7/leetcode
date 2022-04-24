package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeTwoLists(t *testing.T) {
	testCases := []struct {
		nums1 []int
		nums2 []int
		want  []int
	}{
		{
			[]int{1, 2, 4},
			[]int{1, 3, 4},
			[]int{1, 1, 2, 3, 4, 4},
		},
		{
			[]int{},
			[]int{},
			[]int{},
		},
		{
			[]int{},
			[]int{0},
			[]int{0},
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:nums1:%v,nums2:%v,", i+1, tc.nums1, tc.nums2), func(t *testing.T) {
			l1 := generateListNode(tc.nums1)
			l2 := generateListNode(tc.nums2)

			output := mergeTwoLists(l1, l2)

			list := generateList(output)
			assert.Equal(t, list, tc.want)
		})
	}
}
