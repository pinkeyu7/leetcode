package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDestinationCity(t *testing.T) {
	testCases := []struct {
		input [][]string
		want  string
	}{
		{
			[][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}},
			"Sao Paulo",
		},
		{
			[][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}},
			"A",
		},
		{
			[][]string{{"A", "Z"}},
			"Z",
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Test Case #%d,Input:input:%v", i+1, tc.input), func(t *testing.T) {
			output := destCity(tc.input)
			assert.Equal(t, tc.want, output)
		})
	}
}
