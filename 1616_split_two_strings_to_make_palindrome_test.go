package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitTwoStringsToMakePalindrome(t *testing.T) {
	testCases := []struct {
		input  string
		target string
		want   bool
	}{
		{
			"x",
			"y",
			true,
		},
		{
			"xbdef",
			"xecab",
			false,
		},
		{
			"ulacfd",
			"jizalu",
			true,
		},
		{
			"abdef",
			"fecab",
			true,
		},
		{
			"pvhmupgqeltozftlmfjjde",
			"yjgpzbezspnnpszebzmhvp",
			true,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%v,target:%v", i+1, tc.input, tc.target), func(t *testing.T) {
			output := checkPalindromeFormation(tc.input, tc.target)
			assert.Equal(t, tc.want, output)
		})
	}
}
