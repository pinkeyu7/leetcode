package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMajorityElement(t *testing.T) {
	testCases := []struct {
		input []int
		want  int
	}{
		{
			[]int{3, 2, 3},
			3,
		},
		{
			[]int{2, 2, 1, 1, 1, 2, 2},
			2,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%v", i+1, tc.input), func(t *testing.T) {
			output := majorityElement(tc.input)
			assert.Equal(t, tc.want, output)
		})
	}
}
