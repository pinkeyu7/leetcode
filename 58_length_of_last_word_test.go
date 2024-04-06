package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLastWord(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{
			"Hello World",
			5,
		},
		{
			"   fly me   to   the moon  ",
			4,
		},
		{
			"luffy is still joyboy",
			6,
		},
		{
			"a",
			1,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%v", i+1, tc.input), func(t *testing.T) {
			output := lengthOfLastWord(tc.input)
			assert.Equal(t, tc.want, output)
		})
	}
}
