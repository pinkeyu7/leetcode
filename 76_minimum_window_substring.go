package main

import (
	"math"
)

func minWindow(s string, t string) string {
	sLength := len(s)
	tLength := len(t)

	if sLength == 0 || tLength == 0 || sLength < tLength {
		return ""
	}

	scoreMap := make([]int, 128)
	start := 0
	end := 0
	minLength := math.MaxInt64
	startIndex := 0

	for _, char := range t {
		scoreMap[char]++
	}

	for end < sLength {
		if scoreMap[s[end]] > 0 {
			tLength--
		}
		scoreMap[s[end]]--
		end++

		for tLength == 0 {
			if (end - start) < minLength {
				startIndex = start
				minLength = end - start
			}
			if scoreMap[s[start]] == 0 {
				tLength++
			}
			scoreMap[s[start]]++
			start++
		}
	}

	if minLength == math.MaxInt64 {
		return ""
	}

	return s[startIndex : startIndex+minLength]
}
