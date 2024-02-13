package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestStringChain(t *testing.T) {
	testCases := []struct {
		input []string
		want  int
	}{
		{
			[]string{"a", "b", "ba", "bca", "bda", "bdca"},
			4,
		},
		{
			[]string{"xbc", "pcxbcf", "xb", "cxbc", "pcxbc"},
			5,
		},
		{
			[]string{"abcd", "dbqca"},
			1,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%v", i+1, tc.input), func(t *testing.T) {
			output := longestStrChain(tc.input)
			assert.Equal(t, tc.want, output)
		})
	}
}
