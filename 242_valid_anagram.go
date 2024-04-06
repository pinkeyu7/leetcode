package main

import "fmt"

func isAnagram(s string, t string) bool {
	s = sortString(s)
	t = sortString(t)

	fmt.Println(s)
	fmt.Println(t)

	if t == s {
		return true
	} else {
		return false
	}
}
