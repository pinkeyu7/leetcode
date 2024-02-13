package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindFirstPalindromicStringInArray(t *testing.T) {
	testCases := []struct {
		input []string
		want  string
	}{
		{
			[]string{"abc", "car", "ada", "racecar", "cool"},
			"ada",
		},
		{
			[]string{"notapalindrome", "racecar"},
			"racecar",
		},
		{
			[]string{"def", "ghi"},
			"",
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%v", i+1, tc.input), func(t *testing.T) {
			output := firstPalindrome(tc.input)
			assert.Equal(t, tc.want, output)
		})
	}
}
