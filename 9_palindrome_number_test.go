package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		input int
		want  bool
	}{
		{
			121,
			true,
		},
		{
			-121,
			false,
		},
		{
			10,
			false,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:%d", i+1, tc.input), func(t *testing.T) {
			output := isPalindrome(tc.input)
			assert.Equal(t, output, tc.want)
		})
	}
}
