package main

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sMap := make(map[byte]byte)
	tMap := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		v, ok := sMap[s[i]]
		if !ok {
			sMap[s[i]] = t[i]
		} else {
			if v != t[i] {
				return false
			}
		}

		v, ok = tMap[t[i]]
		if !ok {
			tMap[t[i]] = s[i]
		} else {
			if v != s[i] {
				return false
			}
		}
	}

	return true
}
