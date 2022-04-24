package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	testCases := []struct {
		input int
		want  int
	}{
		{
			123,
			321,
		},
		{
			-123,
			-321,
		},
		{
			120,
			21,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:%d", i+1, tc.input), func(t *testing.T) {
			output := reverse(tc.input)
			assert.Equal(t, output, tc.want)
		})
	}
}
