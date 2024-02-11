package main

func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	stack := make([]int, 0)
	indexStack := make([]int, 0)
	cursor := 0

	for i := 0; i < len(temperatures); i++ {
		if cursor == 0 || temperatures[i] <= stack[cursor-1] {
			// Push value
			stack = append(stack, temperatures[i])
			indexStack = append(indexStack, i)
			cursor++
		} else {
			// Pop value
			for j := cursor - 1; j >= 0; j-- {
				if temperatures[i] > stack[j] {
					result[indexStack[j]] = i - indexStack[j]
					cursor--
				} else {
					break
				}
			}
			stack = stack[:cursor]
			indexStack = indexStack[:cursor]

			// Push value
			stack = append(stack, temperatures[i])
			indexStack = append(indexStack, i)
			cursor++
		}
	}

	return result
}
