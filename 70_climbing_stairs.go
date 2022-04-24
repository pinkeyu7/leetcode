package main

func climbStairs(n int) int {
	mapTable := make(map[int]int)
	mapTable[1] = 1
	mapTable[2] = 2

	return climbStairsCallback(n, mapTable)
}

func climbStairsCallback(n int, mapTable map[int]int) int {
	if mapTable[n] == 0 {
		mapTable[n] = climbStairsCallback(n-1, mapTable) + climbStairsCallback(n-2, mapTable)
	}
	return mapTable[n]
}
