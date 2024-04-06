package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsIsomorphic(t *testing.T) {
	testCases := []struct {
		input  string
		target string
		want   bool
	}{
		{
			"egg",
			"add",
			true,
		},
		{
			"foo",
			"bar",
			false,
		},
		{
			"paper",
			"title",
			true,
		},
		{
			"badc",
			"baba",
			false,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%v,target:%v", i+1, tc.input, tc.target), func(t *testing.T) {
			output := isIsomorphic(tc.input, tc.target)
			assert.Equal(t, tc.want, output)
		})
	}
}
