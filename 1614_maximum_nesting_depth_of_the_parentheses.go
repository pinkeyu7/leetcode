package main

func maxDepth(s string) int {
	cursor := 0
	parenthesesStack := make([]string, 0)
	maxDepthCount := 0

	for cursor < len(s) {
		if string(s[cursor]) == "(" {
			parenthesesStack = append(parenthesesStack, "(")
			cursor++
		} else if string(s[cursor]) == ")" {
			if len(parenthesesStack) > maxDepthCount {
				maxDepthCount = len(parenthesesStack)
			}
			parenthesesStack = parenthesesStack[:len(parenthesesStack)-1]
			cursor++
		} else {
			cursor++
		}
	}

	return maxDepthCount
}
