package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordSearchExist(t *testing.T) {
	var testCases = []struct {
		input  [][]byte
		target string
		want   bool
	}{
		{
			[][]byte{[]byte("ABCE"), []byte("SFCS"), []byte("ADEE")},
			"ABCCED",
			true,
		},
		{
			[][]byte{[]byte("ABCE"), []byte("SFCS"), []byte("ADEE")},
			"SEE",
			true,
		},
		{
			[][]byte{[]byte("ABCE"), []byte("SFCS"), []byte("ADEE")},
			"ABCB",
			false,
		},
	}
	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:%v", i+1, tc.input), func(t *testing.T) {
			output := wordSearchExist(tc.input, tc.target)
			assert.Equal(t, tc.want, output)
		})
	}
}
