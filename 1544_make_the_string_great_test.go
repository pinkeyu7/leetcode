package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeGood(t *testing.T) {
	testCases := []struct {
		input string
		want  string
	}{
		{
			"leEeetcode",
			"leetcode",
		},
		{
			"abBAcC",
			"",
		},
		{
			"s",
			"s",
		},
		{
			"Pp",
			"",
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:%v", i+1, tc.input), func(t *testing.T) {
			output := makeGood(tc.input)
			assert.Equal(t, tc.want, output)
		})
	}
}
