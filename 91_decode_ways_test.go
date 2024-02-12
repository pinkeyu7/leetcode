package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeWays(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{
			"12",
			2,
		},
		{
			"226",
			3,
		},
		{
			"06",
			0,
		},
		{
			"10",
			1,
		},
		{
			"1201234",
			3,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%v", i+1, tc.input), func(t *testing.T) {
			output := numDecodings(tc.input)
			assert.Equal(t, tc.want, output)
		})
	}
}
