package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateParenthesis(t *testing.T) {
	testCases := []struct {
		input int
		want  []string
	}{
		{
			3,
			[]string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			1,
			[]string{"()"},
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:%v", i+1, tc.input), func(t *testing.T) {
			output := generateParenthesis(tc.input)
			assert.Equal(t, tc.want, output)
		})
	}
}
