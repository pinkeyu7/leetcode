package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPercentageOfLetterInString(t *testing.T) {
	testCases := []struct {
		input  string
		target byte
		want   int
	}{
		{
			"foobar",
			[]byte("o")[0],
			33,
		},
		{
			"jjjj",
			[]byte("k")[0],
			0,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%v,target:%v", i+1, tc.input, tc.target), func(t *testing.T) {
			output := percentageLetter(tc.input, tc.target)
			assert.Equal(t, tc.want, output)
		})
	}
}
