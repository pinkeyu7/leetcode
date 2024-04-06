package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxDepth(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{
			"(1+(2*3)+((8)/4))+1",
			3,
		},
		{
			"(1)+((2))+(((3)))",
			3,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:%v", i+1, tc.input), func(t *testing.T) {
			output := maxDepth(tc.input)
			assert.Equal(t, tc.want, output)
		})
	}
}
