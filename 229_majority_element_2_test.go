package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMajorityElement2(t *testing.T) {
	testCases := []struct {
		input []int
		want  []int
	}{
		{
			[]int{3, 2, 3},
			[]int{3},
		},
		{
			[]int{1},
			[]int{1},
		},
		{
			[]int{1, 2},
			[]int{1, 2},
		},
		{
			[]int{1, 2, 3},
			[]int{},
		},
		{
			[]int{3, 2, 2, 2, 3},
			[]int{2, 3},
		}}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%v", i+1, tc.input), func(t *testing.T) {
			output := majorityElement2(tc.input)
			assert.Equal(t, tc.want, output)
		})
	}
}
