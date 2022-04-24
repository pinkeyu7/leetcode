package main

import "strconv"

func reverse(x int) int {
	if x >= intMax || x < intMin {
		return 0
	}

	numStr := strconv.Itoa(x)
	if x < 0 {
		numStr = numStr[1:]
	}

	rns := []rune(numStr)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}

	res, _ := strconv.Atoi(string(rns))
	if x < 0 {
		res *= -1
	}

	if res >= intMax || res < intMin {
		return 0
	}

	return res
}
