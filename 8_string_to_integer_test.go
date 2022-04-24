package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyAtoi(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{
			"42",
			42,
		},
		{
			"   -42",
			-42,
		},
		{
			"4193 with words",
			4193,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:%s", i+1, tc.input), func(t *testing.T) {
			output := myAtoi(tc.input)
			assert.Equal(t, output, tc.want)
		})
	}
}
