package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	testCases := []struct {
		input string
		want  bool
	}{
		{
			"()",
			true,
		},
		{
			"()[]{}",
			true,
		},
		{
			"(]",
			false,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%v", i+1, tc.input), func(t *testing.T) {
			output := isValid(tc.input)
			assert.Equal(t, tc.want, output)
		})
	}
}
