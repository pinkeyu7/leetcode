package main

func isValid(s string) bool {
	sStack := make([]rune, 0)

	for _, r := range s {
		if r == 40 || r == 91 || r == 123 {
			sStack = append(sStack, r)
		} else {
			if len(sStack) == 0 {
				return false
			} else if (r == 41 && sStack[len(sStack)-1] == 40) || (r == 93 && sStack[len(sStack)-1] == 91) || (r == 125 && sStack[len(sStack)-1] == 123) {
				sStack = sStack[:len(sStack)-1]
			} else {
				return false
			}
		}
	}

	if len(sStack) > 0 {
		return false
	}

	return true
}
