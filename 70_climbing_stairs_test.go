package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClimbStairs(t *testing.T) {
	testCases := []struct {
		input int
		want  int
	}{
		{
			2,
			2,
		},
		{
			3,
			3,
		},
		{
			5,
			8,
		},
		{
			6,
			13,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%d", i+1, tc.input), func(t *testing.T) {
			output := climbStairs(tc.input)
			assert.Equal(t, tc.want, output)
		})
	}
}
