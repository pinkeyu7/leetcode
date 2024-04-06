package main

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	strMap := map[string][]string{}
	for _, str := range strs {
		res := sortString(str)
		strMap[res] = append(strMap[res], str)
	}

	res := make([][]string, 0)
	for _, value := range strMap {
		sort.Strings(value)
		res = append(res, value)
	}

	return res
}
