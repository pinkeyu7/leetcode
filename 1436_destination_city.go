package main

func destCity(paths [][]string) string {
	scoreMap := make(map[string]string)
	start := paths[0][0]

	for _, path := range paths {
		scoreMap[path[0]] = path[1]
	}

	for scoreMap[start] != "" {
		start = scoreMap[start]
	}

	return start
}
