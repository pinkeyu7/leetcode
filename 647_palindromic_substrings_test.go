package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPalindromicSubstrings(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{
			"abc",
			3,
		},
		{
			"aaa",
			6,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%v", i+1, tc.input), func(t *testing.T) {
			output := countSubstrings(tc.input)
			assert.Equal(t, tc.want, output)
		})
	}
}
