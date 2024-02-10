package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstUniqueCharacter(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{
			"leetcode",
			0,
		},
		{
			"loveleetcode",
			2,
		},
		{
			"aabb",
			-1,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%v", i+1, tc.input), func(t *testing.T) {
			output := firstUniqChar(tc.input)
			assert.Equal(t, tc.want, output)
		})
	}
}
