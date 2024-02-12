package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegerBreak(t *testing.T) {
	testCases := []struct {
		input int
		want  int
	}{
		{
			2,
			1,
		},
		{
			3,
			2,
		},
		{
			4,
			4,
		},
		{
			5,
			6,
		},
		{
			6,
			9,
		},
		{
			7,
			12,
		},
		{
			8,
			18,
		},
		{
			9,
			27,
		},
		{
			10,
			36,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%v", i+1, tc.input), func(t *testing.T) {
			output := integerBreak(tc.input)
			assert.Equal(t, tc.want, output)
		})
	}
}
