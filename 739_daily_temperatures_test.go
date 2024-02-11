package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDailyTemperatures(t *testing.T) {
	testCases := []struct {
		input []int
		want  []int
	}{
		{
			[]int{73, 74, 75, 71, 69, 72, 76, 73},
			[]int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{
			[]int{30, 40, 50, 60},
			[]int{1, 1, 1, 0},
		},
		{
			[]int{30, 60, 90},
			[]int{1, 1, 0},
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%v", i+1, tc.input), func(t *testing.T) {
			output := dailyTemperatures(tc.input)
			assert.Equal(t, tc.want, output)
		})
	}
}
