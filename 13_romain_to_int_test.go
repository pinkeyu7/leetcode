package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRomanToInt(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{
			"III",
			3,
		},
		{
			"IV",
			4,
		},
		{
			"MCMXCIV",
			1994,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:%v", i+1, tc.input), func(t *testing.T) {
			output := romanToInt(tc.input)
			assert.Equal(t, tc.want, output)
		})
	}
}
