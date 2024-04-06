package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinRemoveToMakeValid(t *testing.T) {
	testCases := []struct {
		input string
		want  string
	}{
		{
			"lee(t(c)o)de)",
			"lee(t(c)o)de",
		},
		{
			"a)b(c)d",
			"ab(c)d",
		},
		{
			"))((",
			"",
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:%v", i+1, tc.input), func(t *testing.T) {
			output := minRemoveToMakeValid(tc.input)
			assert.Equal(t, tc.want, output)
		})
	}
}
