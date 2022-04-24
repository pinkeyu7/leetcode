package main

import "sort"

func longestCommonPrefix(strs []string) string {
	sort.SliceStable(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})

	res := strs[0]
	for _, str := range strs {
		for i := 0; i < len(str) && i < len(res); i++ {
			if res[i] != str[i] {
				res = res[0:i]
			}
		}
	}

	return res
}
