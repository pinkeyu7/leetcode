package main

import (
	"strconv"
)

func sequentialDigits(low int, high int) []int {
	lowDigit := len(strconv.Itoa(low))
	highDigit := len(strconv.Itoa(high))

	result := make([]int, 0)

	for i := lowDigit; i <= highDigit; i++ {
		j := uint8(49)
		for j > 48 && j < 58 {
			tempResult := ""
			tempFlag := j
			for k := 0; k < i && tempFlag > 48; k++ {
				tempResult += string(tempFlag)
				tempFlag--
			}

			if len(tempResult) == i {
				temp, _ := strconv.Atoi(reverseString(tempResult))
				if temp >= low && temp <= high {
					result = append(result, temp)
				}
			}
			j++
		}

	}

	return result
}
