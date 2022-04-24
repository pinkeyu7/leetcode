package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddBinary(t *testing.T) {
	testCases := []struct {
		m    string
		n    string
		want string
	}{
		{
			"11",
			"1",
			"100",
		},
		{
			"1010",
			"1011",
			"10101",
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:m:%s,n:%s,", i+1, tc.m, tc.n), func(t *testing.T) {
			output := addBinary(tc.m, tc.n)
			assert.Equal(t, tc.want, output)
		})
	}
}
