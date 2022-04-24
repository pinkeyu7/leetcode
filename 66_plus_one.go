package main

func plusOne(digits []int) []int {
	digitsLen := len(digits)
	carry := 1

	for i := digitsLen - 1; i >= 0; i-- {
		temp := digits[i] + carry
		carry = temp / 10
		digits[i] = temp % 10
	}

	if carry == 0 {
		return digits
	}

	return append([]int{1}, digits...)
}
