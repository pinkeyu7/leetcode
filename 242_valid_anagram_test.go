package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAnagram(t *testing.T) {
	testCases := []struct {
		source string
		target string
		want   bool
	}{
		{
			"anagram",
			"nagaram",
			true,
		},
		{
			"rat",
			"car",
			false,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:source:%v,target:%v", i+1, tc.source, tc.target), func(t *testing.T) {
			output := isAnagram(tc.source, tc.target)
			assert.Equal(t, output, tc.want)
		})
	}
}
