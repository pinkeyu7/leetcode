package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniquePaths(t *testing.T) {
	testCases := []struct {
		m    int
		n    int
		want int
	}{
		{
			3,
			7,
			28,
		},
		{
			3,
			2,
			3,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:m:%d,n:%d,", i+1, tc.m, tc.n), func(t *testing.T) {
			output := uniquePaths(tc.m, tc.n)
			assert.Equal(t, tc.want, output)
		})
	}
}
