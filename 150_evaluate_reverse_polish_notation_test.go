package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvaluateReversePolishNotation(t *testing.T) {
	testCases := []struct {
		input []string
		want  int
	}{
		{
			[]string{"2", "1", "+", "3", "*"},
			9,
		},
		{
			[]string{"4", "13", "5", "/", "+"},
			6,
		},
		{
			[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			22,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%v", i+1, tc.input), func(t *testing.T) {
			output := evalRPN(tc.input)
			assert.Equal(t, tc.want, output)
		})
	}
}
