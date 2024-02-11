package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSequentialDigits(t *testing.T) {
	testCases := []struct {
		input  int
		target int
		want   []int
	}{
		{
			100,
			300,
			[]int{123, 234},
		},
		{
			1000,
			13000,
			[]int{1234, 2345, 3456, 4567, 5678, 6789, 12345},
		},
		{
			10,
			13000,
			[]int{12, 23, 34, 45, 56, 67, 78, 89, 123, 234, 345, 456, 567, 678, 789, 1234, 2345, 3456, 4567, 5678, 6789, 12345},
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%v,traget:%v", i+1, tc.input, tc.target), func(t *testing.T) {
			output := sequentialDigits(tc.input, tc.target)
			assert.Equal(t, tc.want, output)
		})
	}
}
