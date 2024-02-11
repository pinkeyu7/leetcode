package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPalindromicSubsequence(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{
			"bbbab",
			4,
		},
		{
			"cbbd",
			2,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%v", i+1, tc.input), func(t *testing.T) {
			output := longestPalindromeSubseq(tc.input)
			assert.Equal(t, tc.want, output)
		})
	}
}
