package main

import (
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	numStr := strconv.Itoa(x)
	for len(numStr) > 0 {
		if numStr[0] != numStr[len(numStr)-1] {
			return false
		}

		if len(numStr) > 2 {
			numStr = numStr[1 : len(numStr)-1]
		} else {
			break
		}
	}

	return true
}
