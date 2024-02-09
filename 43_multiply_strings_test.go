package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiply(t *testing.T) {
	testCases := []struct {
		num1 string
		num2 string
		want string
	}{
		{
			"0",
			"0",
			"0",
		},
		{
			"3",
			"7",
			"21",
		},
		{
			"13",
			"12",
			"156",
		},
		{
			"498828660196",
			"840477629533",
			"419254329864656431168468",
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:num1:%s,num2:%s,", i+1, tc.num1, tc.num2), func(t *testing.T) {
			output := multiply(tc.num1, tc.num2)
			assert.Equal(t, tc.want, output)
		})
	}
}
