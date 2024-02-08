package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntToRoman(t *testing.T) {
	testCases := []struct {
		input int
		want  string
	}{
		{
			3,
			"III",
		},
		{
			4,
			"IV",
		},
		{
			1994,
			"MCMXCIV",
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:%v", i+1, tc.input), func(t *testing.T) {
			output := intToRoman(tc.input)
			assert.Equal(t, tc.want, output)
		})
	}
}
