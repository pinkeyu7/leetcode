package main

import (
	"strconv"
)

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	num1Len := len(num1)
	num2Len := len(num2)
	result := make([]int, num1Len+num2Len+1)
	for i := num1Len - 1; i >= 0; i-- {
		for j := num2Len - 1; j >= 0; j-- {
			numPadding := num1Len + num2Len - i - j - 2
			number1, _ := strconv.Atoi(string(num1[i]))
			number2, _ := strconv.Atoi(string(num2[j]))
			result[numPadding] += number1 * number2
		}
	}

	for i := 0; i < len(result)-1; i++ {
		result[i+1] += result[i] / 10
		result[i] = result[i] % 10
	}

	resultString := ""
	for i := 0; i < len(result); i++ {
		resultString = strconv.Itoa(result[i]) + resultString
	}

	for i := 0; i < len(resultString); i++ {
		if string(resultString[i]) != "0" {
			resultString = resultString[i:]
			break
		}
	}

	return resultString
}
