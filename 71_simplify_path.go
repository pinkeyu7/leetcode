package main

import "strings"

func simplifyPath(path string) string {
	flag := 0
	res := ""

	pathArgs := strings.Split(path, "/")
	for i := len(pathArgs) - 1; i >= 0; i-- {
		arg := pathArgs[i]
		if arg == "" || arg == "." {
			continue
		}

		if arg == ".." {
			flag++
			continue
		}
		if flag > 0 {
			flag--
			continue
		}

		res = "/" + arg + res
	}

	if len(res) == 0 {
		return "/"
	}

	return res
}
