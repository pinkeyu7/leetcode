package main

func minRemoveToMakeValid(s string) string {
	s, amount := minRemoveToMakeValidSub(s, "(", ")")
	if amount > 0 {
		s = reverseString(s)
		s, _ = minRemoveToMakeValidSub(s, ")", "(")
		s = reverseString(s)
	}

	return s
}

func minRemoveToMakeValidSub(s string, left string, right string) (string, int) {
	cursor := 0
	parenthesesStack := make([]string, 0)

	for cursor < len(s) {
		if string(s[cursor]) == left {
			parenthesesStack = append(parenthesesStack, left)
			cursor++
		} else if string(s[cursor]) == right {
			if len(parenthesesStack) > 0 && string(parenthesesStack[len(parenthesesStack)-1]) == left {
				parenthesesStack = parenthesesStack[:len(parenthesesStack)-1]
				cursor++
			} else {
				s = s[:cursor] + s[cursor+1:]
			}
		} else {
			cursor++
		}
	}
	return s, len(parenthesesStack)
}
