package main

import (
	"sort"
	"strings"
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

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
