package main

func letterCombinations(digits string) []string {
	mapStr := make(map[string][]string)
	mapStr["2"] = []string{"a", "b", "c"}
	mapStr["3"] = []string{"d", "e", "f"}
	mapStr["4"] = []string{"g", "h", "i"}
	mapStr["5"] = []string{"j", "k", "l"}
	mapStr["6"] = []string{"m", "n", "o"}
	mapStr["7"] = []string{"p", "q", "r", "s"}
	mapStr["8"] = []string{"t", "u", "v"}
	mapStr["9"] = []string{"w", "x", "y", "z"}

	return letterCombinationsRecursive(mapStr, digits)
}

func letterCombinationsRecursive(mapStr map[string][]string, digits string) []string {
	var res []string
	if len(digits) == 0 {
		return res
	} else if len(digits) == 1 {
		return mapStr[digits]
	} else {
		callback := letterCombinations(digits[1:])
		for _, char := range mapStr[string(digits[0])] {
			for _, pattern := range callback {
				res = append(res, char+pattern)
			}
		}
	}
	return res
}
