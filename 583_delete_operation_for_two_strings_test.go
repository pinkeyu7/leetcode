package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteOperationForTwoStrings(t *testing.T) {
	testCases := []struct {
		input  string
		target string
		want   int
	}{
		{
			"sea",
			"eat",
			2,
		},
		{
			"leetcode",
			"etco",
			4,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%v,traget:%v", i+1, tc.input, tc.target), func(t *testing.T) {
			output := minDistance(tc.input, tc.target)
			assert.Equal(t, tc.want, output)
		})
	}
}
