package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidSudoku(t *testing.T) {
	testCases := []struct {
		nums [][]byte
		want bool
	}{
		{
			[][]byte{
				[]byte("53..7...."),
				[]byte("6..195..."),
				[]byte(".98....6."),
				[]byte("8...6...3"),
				[]byte("4..8.3..1"),
				[]byte("7...2...6"),
				[]byte(".6....28."),
				[]byte("...419..5"),
				[]byte("....8..79"),
			},
			true,
		},
		{
			[][]byte{
				[]byte("83..7...."),
				[]byte("6..195..."),
				[]byte(".98....6."),
				[]byte("8...6...3"),
				[]byte("4..8.3..1"),
				[]byte("7...2...6"),
				[]byte(".6....28."),
				[]byte("...419..5"),
				[]byte("....8..79"),
			},
			false,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:nums:%v,", i+1, tc.nums), func(t *testing.T) {
			output := isValidSudoku(tc.nums)
			assert.Equal(t, output, tc.want)
		})
	}
}
