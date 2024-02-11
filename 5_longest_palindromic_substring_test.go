package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPalindromicSubstring(t *testing.T) {
	testCases := []struct {
		input string
		want  string
	}{
		{
			"babad",
			"bab",
		},
		{
			"cbbd",
			"bb",
		},
		{
			"caba",
			"aba",
		},
		{
			"aacabdkacaa",
			"aca",
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%v", i+1, tc.input), func(t *testing.T) {
			output := longestPalindrome(tc.input)
			assert.Equal(t, tc.want, output)
		})
	}
}
