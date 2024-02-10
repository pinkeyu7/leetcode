package main

import (
	"sort"
)

func frequencySort(s string) string {
	result := ""
	charMap := make(map[string]int)

	for _, c := range s {
		charMap[string(c)]++
	}

	keys := make([]string, 0)
	for key := range charMap {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return charMap[keys[i]] > charMap[keys[j]]
	})

	for _, key := range keys {
		for i := 0; i < charMap[key]; i++ {
			result += key
		}
	}

	return result
}
