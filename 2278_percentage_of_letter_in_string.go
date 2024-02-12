package main

func percentageLetter(s string, letter byte) int {
	scoreMap := make([]int, 128)
	for _, char := range s {
		scoreMap[char]++
	}

	return scoreMap[letter] * 100 / len(s)
}
