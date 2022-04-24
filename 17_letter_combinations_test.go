package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLetterCombinations(t *testing.T) {
	testCases := []struct {
		input string
		want  []string
	}{
		{
			"23",
			[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			"",
			nil,
		},
		{
			"2",
			[]string{"a", "b", "c"},
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:%v", i+1, tc.input), func(t *testing.T) {
			output := letterCombinations(tc.input)
			assert.Equal(t, output, tc.want)
		})
	}
}
