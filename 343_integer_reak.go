package main

func integerBreak(n int) int {
	if n == 2 {
		return 1
	} else if n == 3 {
		return 2
	} else if n == 4 {
		return 4
	} else {
		quotient := n / 3
		remainder := n % 3

		result := 1
		if remainder == 0 {
			for i := 0; i < quotient; i++ {
				result *= 3
			}
			return result
		} else if remainder == 1 {
			for i := 0; i < quotient-1; i++ {
				result *= 3
			}
			return result * 4
		} else {
			for i := 0; i < quotient; i++ {
				result *= 3
			}
			return result * 2
		}
	}
}
