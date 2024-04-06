package main

func makeGood(s string) string {
	return makeGoodRecursive(s)
}

func makeGoodRecursive(s string) string {
	cursor := 0
	needNextRound := false
	for cursor < len(s)-1 {
		if s[cursor]-32 == s[cursor+1] || s[cursor]+32 == s[cursor+1] {
			s = s[:cursor] + s[cursor+2:]
			needNextRound = true
		} else {
			cursor++
		}
	}

	if needNextRound {
		return makeGoodRecursive(s)
	} else {
		return s
	}
}
