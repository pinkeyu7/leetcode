package main

import "strconv"

func myAtoi(s string) int {
	strLen := len(s)
	resStr := ""
	flag := false

	for i := 0; i < strLen; i++ {
		if !flag && (s[i] == 43 || s[i] == 45) {
			resStr += string(s[i])
			flag = true
		} else if s[i] > 47 && s[i] < 58 {
			resStr += string(s[i])
			flag = true
		} else {
			if s[i] == 32 {
				if resStr == "" {
					continue
				} else {
					break
				}
			} else {
				if resStr == "" {
					return 0
				} else {
					break
				}
			}
		}
	}

	res, _ := strconv.Atoi(resStr)

	if res < intMin {
		return intMin
	}

	if res > intMax {
		return intMax
	}

	return res
}
