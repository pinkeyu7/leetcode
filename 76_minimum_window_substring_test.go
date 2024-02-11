package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumWindowSubstring(t *testing.T) {
	testCases := []struct {
		input  string
		target string
		want   string
	}{
		{
			"ADOBECODEBANC",
			"ABC",
			"BANC",
		},
		{
			"a",
			"a",
			"a",
		},
		{
			"a",
			"aa",
			"",
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%v,traget:%v", i+1, tc.input, tc.target), func(t *testing.T) {
			output := minWindow(tc.input, tc.target)
			assert.Equal(t, tc.want, output)
		})
	}
}
