package main

func isInterleave(s1 string, s2 string, s3 string) bool {
	mapVar := map[[3]int]bool{}
	mapVar[[3]int{len(s1), len(s2), len(s3)}] = true

	return isInterleaveCallback(s1, s2, s3, 0, 0, 0, mapVar)
}

func isInterleaveCallback(s1, s2, s3 string, i1, i2, i3 int, mapVar map[[3]int]bool) bool {
	ok, checked := mapVar[[3]int{i1, i2, i3}]
	if checked {
		return ok
	}

	flag := false
	if i1 < len(s1) && s1[i1] == s3[i3] {
		flag = isInterleaveCallback(s1, s2, s3, i1+1, i2, i3+1, mapVar)
	}
	if i2 < len(s2) && s2[i2] == s3[i3] {
		flag = flag || isInterleaveCallback(s1, s2, s3, i1, i2+1, i3+1, mapVar)
	}
	mapVar[[3]int{i1, i2, i3}] = flag
	return flag
}
