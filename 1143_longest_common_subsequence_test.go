package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonSubsequence(t *testing.T) {
	testCases := []struct {
		input  string
		target string
		want   int
	}{
		{
			"abcde",
			"ace",
			3,
		},
		{
			"abc",
			"abc",
			3,
		},
		{
			"abc",
			"def",
			0,
		},
		{
			"bl",
			"yby",
			1,
		},
		{
			"psnw",
			"vozsh",
			1,
		},
		{
			"ezupkr",
			"ubmrapg",
			2,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%v,traget:%v", i+1, tc.input, tc.target), func(t *testing.T) {
			output := longestCommonSubsequence(tc.input, tc.target)
			assert.Equal(t, tc.want, output)
		})
	}
}
