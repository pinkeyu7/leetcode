package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumSquares(t *testing.T) {
	testCases := []struct {
		input int
		want  int
	}{
		{
			12,
			3,
		},
		{
			13,
			2,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%v", i+1, tc.input), func(t *testing.T) {
			output := numSquares(tc.input)
			assert.Equal(t, tc.want, output)
		})
	}
}
