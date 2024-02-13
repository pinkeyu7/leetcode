package main

import (
	"sort"
)

func longestStrChain(words []string) int {
	scoreMap := make(map[string]int)

	sort.SliceStable(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	ans := 0
	for _, word := range words {
		scoreMap[word] = 1
		for i := 0; i < len(word); i++ {
			tempString := word[:i] + word[i+1:]
			if scoreMap[tempString] > 0 {
				if scoreMap[tempString] >= scoreMap[word] {
					scoreMap[word] = scoreMap[tempString] + 1
				}
			}
		}
		if scoreMap[word] > ans {
			ans = scoreMap[word]
		}
	}

	return ans
}
