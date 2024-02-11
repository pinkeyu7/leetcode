package main

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	cursor := 0
	result := 0
	init := true

	for _, token := range tokens {
		if token == "+" {
			if cursor == 1 {
				num := stack[cursor-1]
				result += num

				cursor--
				stack = stack[:cursor]
			} else {
				num1 := stack[cursor-2]
				num2 := stack[cursor-1]

				cursor -= 2
				stack = stack[:cursor]

				cursor++
				stack = append(stack, num1+num2)
			}
		} else if token == "-" {
			if cursor == 1 {
				num := stack[cursor-1]
				result -= num

				cursor--
				stack = stack[:cursor]
			} else {
				num1 := stack[cursor-2]
				num2 := stack[cursor-1]

				cursor -= 2
				stack = stack[:cursor]

				cursor++
				stack = append(stack, num1-num2)
			}
		} else if token == "*" {
			if cursor == 1 {
				num := stack[cursor-1]
				result *= num

				cursor--
				stack = stack[:cursor]
			} else {
				num1 := stack[cursor-2]
				num2 := stack[cursor-1]

				cursor -= 2
				stack = stack[:cursor]

				cursor++
				stack = append(stack, num1*num2)
			}
		} else if token == "/" {
			if cursor == 1 {
				num := stack[cursor-1]
				result /= num

				cursor--
				stack = stack[:cursor]
			} else {
				num1 := stack[cursor-2]
				num2 := stack[cursor-1]

				cursor -= 2
				stack = stack[:cursor]

				cursor++
				stack = append(stack, num1/num2)
			}
		} else {
			num, _ := strconv.Atoi(token)
			if init {
				result = num
				init = false
			} else {
				stack = append(stack, num)
				cursor++
			}
		}
	}
	return result
}
