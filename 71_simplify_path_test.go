package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimplifyPath(t *testing.T) {
	testCases := []struct {
		input string
		want  string
	}{
		{
			"/home/",
			"/home",
		},
		{
			"/../",
			"/",
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%s", i+1, tc.input), func(t *testing.T) {
			output := simplifyPath(tc.input)
			assert.Equal(t, tc.want, output)
		})
	}
}
