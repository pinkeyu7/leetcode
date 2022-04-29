package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsInterleave(t *testing.T) {
	testCases := []struct {
		s1   string
		s2   string
		s3   string
		want bool
	}{
		{
			"aabcc",
			"dbbca",
			"aadbbcbcac",
			true,
		},
		{
			"aabcc",
			"dbbca",
			"aadbbbaccc",
			false,
		},
		{
			"",
			"",
			"",
			true,
		},
		{
			"abababababababababababababababababababababababababababababababababababababababababababababababababbb",
			"babababababababababababababababababababababababababababababababababababababababababababababababaaaba",
			"abababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababbb",
			false,
		},
		{
			"ab",
			"a",
			"c",
			false,
		},
		{
			"",
			"a",
			"c",
			false,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:s1:%s,s2:%s,s3:%s", i+1, tc.s1, tc.s2, tc.s3), func(t *testing.T) {
			output := isInterleave(tc.s1, tc.s2, tc.s3)
			assert.Equal(t, tc.want, output)
		})
	}
}
